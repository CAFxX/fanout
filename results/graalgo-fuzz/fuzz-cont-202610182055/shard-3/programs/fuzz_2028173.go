package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint32 }
func genFn0() int {
    var v6 int8 = 45
    var v7 string = "a"
    switch ((v7 + v7) + ((v7 + v7) + (v7 + v7))) {
    case "":
        v6 -= (v6 % ((v6) | 1))
    case "foo bar":
        switch 128 {
        case 787:
            var v8 int32 = 89
            _ = v8
        case 892:
            if (2 != 128) {
                var v9 uint32 = 417
                var v10 uint8 = 42
                v6 += (v6 ^ 1)
                _ = v10
                _ = v9
            }
        case -418:
            v7 = (("a" + "a") + ("hello" + v7))
        default:
            v11 := GenS1{F0: 2, F1: v7}
            _ = v11
        }
    case "hello":
        v12 := GenS1{F0: 883, F1: ""}
        _ = v12
    default:
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := []int{}
            fmt.Println(v7, fi13)
            _ = fi13
            _ = v14
        }
    }
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    var v3 bool = (4294967295 >= 237)
    var v4 bool = v3
    var v5 uint = 255
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS1{F0: 32767, F1: "a"}
        v3 = (0 <= 42)
        v5 += v5
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v5 -= v5
        _ = fi17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        _ = len(v2)
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS1{F0: 42, F1: "foo bar"}
            var v21 int8 = 42
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 int = (fi18 & (fi18 | fi18))
        _ = fi18
        _ = v22
    }
    if (!((0.1 - 3.14159) >= 0.0)) {
        if true {
            fmt.Println(v3)
            var v23 uint32 = 2147483647
            _ = v23
        } else {
            switch -2147483648 {
            case -501:
                var v24 int = 0
                _ = v24
            case 128:
                fmt.Println(v5)
            default:
                var v25 int16 = 1
                _ = v25
            }
            var v26 uint16 = 2
            _ = v26
        }
        switch 529 {
        case -9223372036854775808:
            var v27 uint16 = 1
            _ = v27
        case -1:
            var v28 uint = v5
            _ = v28
        case 127:
            v4 = (1.5 < 3.14159)
        }
    }
    v5 -= v5
    if (60 != 80) {
        _ = len(v2)
        fmt.Println(v4, v3, v5)
        switch "hello" {
        case "hello":
            v5 = 65535
        default:
            switch 42 {
            case 1:
                switch "foo bar" {
                case "":
                    v5 = (v5 & (^(v5)))
                default:
                    var v29 uint = 4294967295
                    _ = v29
                }
            case 42:
                v3 = (v5 <= v5)
            case 151:
                var v30 int8 = -85
                _ = v30
            default:
                for fi31 := 0; fi31 < 19; fi31++ {
                    delete(v1, genFn0())
                    _ = fi31
                }
            }
        }
    } else {
        v32 := GenS2{F0: 171}
        _ = v32
    }
    var v33 int16 = -41
    if (!(2 == 127)) {
        if ((1 < 0) && (v5 != v5)) {
            v4 = (2 <= 255)
            v34 := GenS2{F0: 65535}
            v3 = false
            _ = v34
        }
        v35 := GenS1{F0: 69, F1: "hello"}
        for fi36 := 0; fi36 < 1; fi36++ {
            v33 -= (-(v33))
            if v4 {
                v2 = append(v2, ("foo bar" + ""))
                v37 := v1[fi36]
                v3 = (false && false)
                _ = v37
            } else {
                if (842 == 1) {
                    var v38 bool = false
                    var v39 float32 = float32(0.1)
                    _ = v38
                    _ = v39
                }
                var v40 uint32 = 42
                _ = v40
            }
            _ = fi36
        }
        _ = v35
    } else {
        if (304 == 42) {
            var v41 int16 = (v33 + v33)
            v5 = (v5 ^ (^(v5)))
            var v42 uint16 = 926
            _ = v41
            _ = v42
        } else {
            switch genFn0() {
            case -1:
                for fi43 := 0; fi43 < 2; fi43++ {
                    v3 = (false || v4)
                    _ = fi43
                }
            default:
                v44 := GenS1{F0: 32767, F1: "hello"}
                _ = v44
            }
        }
        for fi45 := 0; fi45 < 9; fi45++ {
            for fi46 := 0; fi46 < 20; fi46++ {
                fmt.Println(fi45, fi45)
                _ = fi46
            }
            var v47 uint64 = 2
            _ = fi45
            _ = v47
        }
    }
    var v48 int64 = 0
    v5 -= v5
    var v49 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
