#include "utils.hpp" // "symbols.hpp" -> iostream, unordered_map, variant; string, fstream, vector
#include <optional>
//todo: Implement method to add a locally scoped Variables table

SymbolTable::SymbolTable(Variables vars) : m_g_vars(vars)
{}

Variables SymbolTable::get_g_vars()
{
    return this->m_g_vars;
}

void SymbolTable::add_var(std::string name, AnyType value)
{
    this->m_g_vars.insert_or_assign(name, value);
}

void SymbolTable::display_vars()
{
    std::cout << "\nCurrent Variables:" << std::endl << std::endl;
    size_t i = this->m_g_vars.size();
    for (auto& var : this->m_g_vars) {
        std::string varType = get_var_type(var.second);
        std::cout << "Name: " << var.first << std::endl;
        std::cout << "Value: " << var.second;
        std::cout << "Type: " << varType << std::endl << std::endl;
        i--;
    }
}

bool SymbolTable::contains_all(std::vector<std::string> found_vars)
{
    for (std::string var : found_vars) {
        if (!contains_key(this->m_g_vars, var)) return false;
    }
    return true;
}

std::optional<AnyType> SymbolTable::get_val(std::string name)
{
    for (auto pair : this->m_g_vars) {
        if (pair.first == name) return pair.second;
    }
    return std::nullopt;
}

void SymbolTable::display_l_vars_list() 
{
    for (auto list : this->m_l_vars) {
        std::cout << list.first << ":" << std::endl;
        for (auto item : list.second) {
            std::cout << item.first << ": " << item.second; 
        }
    }
}

std::optional<Variables> SymbolTable::get_l_vars_list(std::string name) 
{
    for (auto list : this->m_l_vars) {
        if (list.first == name) {
            return list.second;
        }
    }
    return std::nullopt;
}

void SymbolTable::add_l_var(std::string name, AnyType val, std::string l_var_name)
{
    for (auto item : this->m_l_vars) {
        if (item.first == l_var_name) {
            this->m_l_vars.at(l_var_name).insert_or_assign(name, val);
            return;
        }
    }
}

void SymbolTable::new_l_vars(std::string name)
{
    std::unordered_map<std::string, AnyType> vars;
    this->m_l_vars.insert_or_assign(name, vars);
}

void SymbolTable::pop_l_vars(std::string vars)
{
    for (auto item : this->m_l_vars) {
        if (item.first == vars) {
            this->m_l_vars.erase(vars);
        }
    }
}
