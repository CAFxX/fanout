package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
func genFn0() bool {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    v9 := []int{2147483647, -9223372036854775808, 42, 452}
    switch -1 {
    case 4294967295:
        var v10 bool = false
        _ = v10
    case -158:
        for fi11 := 0; fi11 < 9; fi11++ {
            if (-128 != 1) {
                fmt.Println(fi11)
            }
            if len(v9) > 0 {
                v9[0] = fi11
            }
            _ = fi11
        }
    default:
        var v12 int = 32767
        _ = v12
    }
    _ = v9
    return false
}

func main() {
    v1 := []string{("foo bar" + ("foo bar" + "\n\t")), ("\n\t" + "a")}
    var v2 uint16 = 0
    var v3 int64 = 1000
    var v4 string = "foo bar"
    var v5 bool = (1.5 > 0.1)
    var v6 float64 = (0.0 * 1e10)
    v13 := GenS1{F0: 42}
    var v14 string = ""
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 string = (v14 + v14)
        var v17 float64 = float64(0.1)
        v18 := GenS2{F0: 1}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    switch -1 {
    case 32767:
        v14 = "a"
    case -2147483648:
        v14 = "\n\t"
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 bool = (32767 < 42)
        var v21 uint16 = (v2 | (v2 * (v2 * v2)))
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: 602}
    if (42 != 399) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v23 bool = v5
        _ = v23
    }
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
