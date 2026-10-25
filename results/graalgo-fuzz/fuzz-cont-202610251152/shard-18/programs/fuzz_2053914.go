package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float64 = 1.5
    var v2 int64 = -9223372036854775808
    var v3 string = ("foo bar" + "\n\t")
    var v4 int64 = (^(v2))
    v5 := make(map[uint32]bool)
    fmt.Println(v4)
    v6 := GenS1{F0: 912}
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 18446744073709551615}
        if (false && (!true)) {
            if ((18446744073709551615 >= 2) || ("hello" >= (v3 + v3))) {
                v3 = v3
            }
            var v9 int = (fi7 & (fi7 ^ (^(fi7))))
            _ = v9
        } else {
            if false {
                v1 += v1
            }
            v10 := GenS2{F0: 217}
            var v11 string = ("a" + (v3 + (v3 + v3)))
            _ = v10
            _ = v11
        }
        var v12 int64 = 0
        _ = fi7
        _ = v12
        _ = v8
    }
    var v13 int = 0
    var v14 int = v13
    if false {
        v15 := GenS1{F0: 65535}
        v16 := GenS2{F0: 123}
        var v17 int16 = 0
        _ = v15
        _ = v16
        _ = v17
    } else {
        var v18 int32 = 1
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 bool = false
        _ = fi19
        _ = v20
    }
    v5[2147483648] = (false || (!(0.0 > (0.1 - 0.0))))
    if true {
        var v21 uint8 = 141
        v1 = float64(-2.25)
        _ = v21
    } else {
        var v22 int = (v14 + -2147483648)
        _ = len(v5)
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 int16 = 42
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    var v25 string = (((v3 + ("foo bar" + ("hello" + v3))) + (v3 + "\n\t")) + "hello")
    if true {
        var v26 uint8 = 1
        _ = v26
    } else {
        for fi27 := 0; fi27 < 16; fi27++ {
            v25 = v3
            _ = fi27
        }
        v25 = (v3 + ("" + (v3 + "")))
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 int32 = 1
            v29 -= (v29 + v29)
            _ = fi28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
