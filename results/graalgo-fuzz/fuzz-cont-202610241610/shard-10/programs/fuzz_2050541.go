package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{"hello"}
    var v3 int64 = v1
    fmt.Println(v1, v1)
    for fi4 := 0; fi4 < 9; fi4++ {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        v2 = append(v2, "\n\t")
        v2 = append(v2, "")
        _ = fi4
    }
    var v5 int32 = -792
    var v6 int8 = 0
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 string = ""
        for fi9 := 0; fi9 < 3; fi9++ {
            if len(v2) > 0 {
                v2[0] = ("a" + v8)
            }
            _ = fi9
        }
        v10 := GenS1{F0: v8}
        _ = fi7
        _ = v10
        _ = v8
    }
    var v11 int8 = (v6 % (((v6 | (v6 & 0))) | 1))
    v12 := GenS1{F0: ""}
    v6 += (v11 / ((v11) | 1))
    var v13 uint16 = 2
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            v5 = v5
            var v16 uint64 = 1
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v5 += 42
    v1 += (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
