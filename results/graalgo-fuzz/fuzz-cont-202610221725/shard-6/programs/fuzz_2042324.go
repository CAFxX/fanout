package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int64 = 1
    var v2 uint8 = 42
    var v3 bool = true
    v2 += (v2 | (v2 >> ((18446744073709551615) & 63)))
    for fi4 := 0; fi4 < 13; fi4++ {
        v3 = (!v3)
        _ = fi4
    }
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            fmt.Println(v1, fi6, fi6)
            if v3 {
                v7 := GenS1{F0: (v2 ^ (v2 % ((v2) | 1)))}
                var v8 uint = 65535
                _ = v7
                _ = v8
            }
            v1 = 1
            _ = fi6
        }
        v3 = true
        v2 = (v2 >> ((42) & 63))
        _ = fi5
    }
    if v3 {
        var v9 int8 = 42
        _ = v9
    } else {
        var v10 int64 = v1
        v3 = (260 == 42)
        fmt.Println(v2, v3)
        _ = v10
    }
    if v3 {
        if false {
            if false {
                fmt.Println(v1, v3, v2)
                for fi11 := 0; fi11 < 7; fi11++ {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        v2 = (v2 - (v2 * 42))
                        _ = fi12
                    }
                    if (!true) {
                        v2 = (v2 ^ (^(v2)))
                        v3 = v3
                    }
                    v13 := []int{}
                    _ = fi11
                    _ = v13
                }
            }
            v14 := []int{-1, -9223372036854775808, 127, 247}
            var v15 int8 = 1
            _ = v14
            _ = v15
        }
        fmt.Println(v1, v3)
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 int8 = 54
            v1 += (v1 % (((v1 << ((42) & 63))) | 1))
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS2{F0: 989}
    var v19 int = 42
    switch ("" + ("hello" + (("\n\t" + "\n\t") + ("" + "")))) {
    case "":
        v19 = (-(v19))
    }
    v20 := GenS1{F0: v2}
    var v21 uint8 = (v2 + (v2 - 0))
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := []int{fi22}
        _ = fi22
        _ = v23
    }
    v2 -= (^(v2))
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
