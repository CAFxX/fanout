package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 bool, p1 int32) bool {
    var v4 int64 = -1
    var v5 float32 = 1.5
    if (4294967295 >= 988) {
        v6 := []int{}
        switch ("\n\t" + (("a" + "a") + "foo bar")) {
        case "a":
            v6 = append(v6, 1)
        case "foo bar":
            var v7 int8 = 0
            _ = v7
        case "\n\t":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = v6
    } else {
        fmt.Println(v4, v4, p1)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (true && true)
}

func main() {
    var v1 int = 4294967295
    var v2 int8 = -117
    var v3 float32 = 3.14159
    var v8 bool = false
    v8 = v8
    v2 = (v2 / ((v2) | 1))
    switch (v1 * v1) {
    case -1:
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := []int{(^(v1))}
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v1, v1, v3)
    fmt.Println(v2, v2)
    switch (v1 ^ (^(v1))) {
    case 32768:
        v11 := []int{0, (v1 >> ((255) & 63))}
        _ = v11
    default:
        for fi12 := 0; fi12 < 20; fi12++ {
            if v8 {
                v13 := GenS1{F0: 771}
                _ = v13
            }
            _ = fi12
        }
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        switch -146 {
        case 106:
            fmt.Println(v3, v8, v8)
        case 2147483647:
            var v15 float32 = v3
            _ = v15
        default:
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := GenS1{F0: 2147483648}
                var v18 int32 = 42
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        var v19 bool = (false && (v8 && (4294967295 >= 4294967295)))
        fmt.Println(v8, v19, fi14)
        _ = fi14
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
