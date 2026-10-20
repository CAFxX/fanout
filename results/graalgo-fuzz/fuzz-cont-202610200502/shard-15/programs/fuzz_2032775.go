package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32) int8 {
    var v4 uint64 = 807
    var v5 uint = 2
    _ = p0
    _ = v4
    _ = v5
    return -103
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("" + "foo bar")
    var v3 int = 128
    var v6 float32 = 0.0
    if true {
        for fi7 := 0; fi7 < 6; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 string = (("" + "foo bar") + v2)
                _ = fi8
                _ = v9
            }
            switch 42 {
            case 440:
                v10 := []int{fi7, 2147483647, (v3 | 42)}
                _ = v10
            }
            v2 = v2
            _ = fi7
        }
        switch v3 {
        case 32767:
            var v11 int8 = -1
            _ = v11
        default:
            if false {
                fmt.Println(v6)
            } else {
                if true {
                    v12 := []int{}
                    v2 = (v2 + (v2 + "a"))
                    _ = v12
                }
                switch (v3 | v3) {
                case 9223372036854775807:
                    if true {
                        v13 := GenS1{F0: 0}
                        for fi14 := 0; fi14 < 8; fi14++ {
                            v3 += v3
                            _ = fi14
                        }
                        _ = v13
                    } else {
                        fmt.Println(v2)
                    }
                }
            }
        }
        v3 -= (^(v3))
    }
    var v15 string = v2
    v16 := v1[("hello" + (v2 + ("a" + v15)))]
    var v17 int64 = 0
    fmt.Println(v3, v17, v3)
    if (-9223372036854775808 < 882) {
        fmt.Println(v6)
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := []int{(^(fi18)), fi18, (v3 % ((v3) | 1)), (^(fi18))}
            _ = fi18
            _ = v19
        }
    } else {
        var v20 uint32 = 4294967295
        v2 = v2
        _ = v20
    }
    var v21 uint8 = 94
    if ((v16 + v15) != v16) {
        v22 := []int{}
        v15 = "foo bar"
        if (2147483647 < 1) {
            v23 := GenS2{F0: 100}
            _ = v23
        }
        _ = v22
    } else {
        v24 := GenS2{F0: 1}
        v25 := GenS1{F0: 18446744073709551615}
        var v26 int = v3
        _ = v24
        _ = v25
        _ = v26
    }
    var v27 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v6
}
