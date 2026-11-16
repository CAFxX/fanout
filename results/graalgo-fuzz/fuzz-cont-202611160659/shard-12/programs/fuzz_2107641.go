package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    v1 := []string{"hello", ("hello" + "\n\t"), "foo bar", ""}
    var v2 bool = true
    var v3 float32 = 3.14159
    v3 += 1.5
    v3 = ((v3 / v3) / v3)
    if ((!false) || (false && (-128 >= 100))) {
        for fi4 := 0; fi4 < 5; fi4++ {
            var v5 int64 = 9223372036854775807
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi4
            _ = v5
        }
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        fmt.Println(v3)
    }
    v6 := GenS1{F0: 32767, F1: -391}
    v7 := GenS2{F0: -32768, F1: 18446744073709551615}
    v2 = v2
    if len(v1) > 0 {
        v1[0] = ("\n\t" + ("foo bar" + "\n\t"))
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v3 += v3
        var v9 bool = (v2 && v2)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 165, F1: -9223372036854775808}
    v2 = (-100 >= 0)
    v2 = (127 == 127)
    v1 = append(v1, "foo bar")
    v2 = (v2 && true)
    var v11 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
