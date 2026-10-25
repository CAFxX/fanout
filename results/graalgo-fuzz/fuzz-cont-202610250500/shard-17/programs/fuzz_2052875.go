package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int8; F1 uint; F2 string }
func genFn0(p0 string) int16 {
    fmt.Println(p0)
    var v7 float64 = 3.14159
    if (32768 == -93) {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint64 = 0
            v10 := []int{fi8}
            for fi11 := 0; fi11 < 5; fi11++ {
                if (1 == 832) {
                    v9 += (v9 % ((v9) | 1))
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v9 -= (v9 >> ((614) & 63))
                _ = fi11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        v12 := GenS1{F0: 0}
        for fi13 := 0; fi13 < 10; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 uint16 = 65535
                v16 := []int{}
                v7 += float64(-2.25)
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi13
        }
        _ = v12
    }
    switch 128 {
    case -9223372036854775808:
        p0 = "foo bar"
    case -2147483648:
        v17 := []int{-1, -414}
        _ = v17
    case 1:
        fmt.Println(p0, v7)
    default:
        v7 -= ((v7 - float64(0.0)) + v7)
    }
    var v18 uint16 = 32767
    _ = p0
    _ = v18
    _ = v7
    return 32767
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = 3.14159
    var v3 int64 = -384
    v4 := []uint32{}
    v5 := make(map[uint32]bool)
    var v6 bool = (2 > 127)
    v1 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v1, v3)
    fmt.Println(v1, v3, v1)
    if (!true) {
        if v6 {
            var v19 int64 = (v3 % ((-1) | 1))
            switch "hello" {
            case "hello":
                var v20 uint16 = 308
                _ = v20
            case "":
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 string = ("\n\t" + "\n\t")
                    _ = fi21
                    _ = v22
                }
            default:
                v4 = append(v4, 646)
            }
            _ = v19
        } else {
            fmt.Println(v6)
            v23 := GenS2{F0: v6, F1: true}
            if ((v3 >> ((42) & 63)) > v3) {
                if true {
                    if true {
                        v2 = v2
                    } else {
                        v1 = v1
                    }
                } else {
                    var v24 uint8 = 164
                    var v25 int16 = -32768
                    var v26 int64 = (v3 - 42)
                    _ = v24
                    _ = v25
                    _ = v26
                }
                switch "hello" {
                case "\n\t":
                    var v27 int8 = 1
                    _ = v27
                case "foo bar":
                    _ = len(v5)
                }
                var v28 int = 2147483647
                _ = v28
            }
            _ = v23
        }
        switch -644 {
        case 499:
            if false {
                v6 = (v6 || v6)
                v1 -= (v1 + v1)
            } else {
                if (v2 >= v2) {
                    var v29 string = (("a" + "a") + ("foo bar" + ("hello" + "")))
                    _ = v29
                }
                if (647 < 1000) {
                    v3 -= v3
                    v6 = (1000 != 760)
                }
            }
        case 4294967295:
            v30 := GenS3{F0: v6, F1: ("foo bar" + "\n\t")}
            _ = v30
        }
    }
    v31 := GenS1{F0: 127}
    if v6 {
        for fi32 := 0; fi32 < 3; fi32++ {
            v33 := GenS3{F0: false, F1: ("\n\t" + "hello")}
            _ = fi32
            _ = v33
        }
        v34 := GenS2{F0: false, F1: (255 >= 2)}
        for fi35 := 0; fi35 < 19; fi35++ {
            switch 4294967295 {
            case -307:
                v3 += v3
            case 127:
                var v36 float32 = (3.14159 / (-2.25 + -2.25))
                _ = v36
            case 281:
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            }
            v37 := GenS4{F0: 255, F1: true}
            _ = fi35
            _ = v37
        }
        _ = v34
    }
    if false {
        v6 = false
    } else {
        v3 = v3
    }
    var v38 int64 = v3
    if ("a" < ("" + (("" + ("hello" + "")) + "foo bar"))) {
        v3 = (v38 % ((v38) | 1))
        _ = len(v4)
    } else {
        v39 := GenS5{F0: -128, F1: 991, F2: "foo bar"}
        _ = v39
    }
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    var v40 bool = true
    v1 -= (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
