package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint }
func genFn0() int16 {
    var v7 float32 = -2.25
    var v8 int64 = 0
    switch 2147483647 {
    case 1:
        v9 := []int{1}
        _ = v9
    case 32768:
        if true {
            if (51 < int8(v8)) {
                switch 32767 {
                case 127:
                    var v10 uint64 = 4294967295
                    _ = v10
                case -2147483648:
                    v11 := []int{}
                    _ = v11
                }
            } else {
                v12 := GenS1{F0: "foo bar"}
                fmt.Println(v8, v7)
                _ = v12
            }
            fmt.Println(v8, v8)
        } else {
            v13 := []int{}
            var v14 uint8 = 21
            var v15 uint16 = 0
            _ = v13
            _ = v14
            _ = v15
        }
    default:
        var v16 float64 = float64(v7)
        _ = v16
    }
    _ = v7
    _ = v8
    return -32768
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]bool)
    var v3 uint8 = 140
    var v4 bool = true
    v5 := make(map[string]uint32)
    var v6 int = int(v1)
    v17 := []int{v6, -9223372036854775808, -9223372036854775808, (v6 | (v6 >> ((v1) & 63)))}
    var v18 int8 = -1
    v4 = ("a" < ((("foo bar" + "") + (("foo bar" + "") + "foo bar")) + ("\n\t" + (("foo bar" + ("foo bar" + ((((("a" + "\n\t") + "\n\t") + "\n\t") + "a") + "\n\t"))) + "a"))))
    fmt.Println(v6)
    v6 = v6
    if len(v17) > 0 {
        v17[0] = (v6 >> (((v1 % (((v1 | uint(v3))) | 1))) & 63))
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            switch -9223372036854775808 {
            case 2147483647:
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            case 128:
                var v21 float32 = -2.25
                _ = v21
            case -390:
                if (((!v4) && (-2147483648 >= -1)) && v4) {
                    if (!false) {
                        v18 = v18
                        var v22 int = (fi20 / (((fi19 >> (((v1 / ((v1) | 1))) & 63))) | 1))
                        v23 := GenS2{F0: 2}
                        _ = v22
                        _ = v23
                    }
                    var v24 uint = (v1 << (((v1 | (v1 & 2))) & 63))
                    _ = v24
                }
            default:
                var v25 uint16 = 9
                _ = v25
            }
            fmt.Println(v1, fi20, v3)
            var v26 float32 = -2.25
            var v27 uint = (v1 >> ((v1) & 63))
            var v28 int64 = 1
            v29 := GenS2{F0: 18446744073709551615}
            _ = fi20
            _ = v26
            _ = v27
            _ = v28
            _ = v29
        }
        fmt.Println(v4)
        for fi30 := 0; fi30 < 14; fi30++ {
            switch ("a" + ("\n\t" + "")) {
            case "a":
                var v31 uint16 = 2
                _ = v31
            case "":
                v32 := GenS3{F0: -2147483648, F1: 1, F2: 65535}
                _ = v32
            }
            v1 += v1
            _ = fi30
        }
        _ = fi19
    }
    v33 := GenS2{F0: 826}
    if true {
        if false {
            v34 := GenS4{F0: (^(v1))}
            if (((!(!true)) && (!(v4 || v4))) || true) {
                v35 := GenS4{F0: (v1 - v1)}
                fmt.Println(v1, v4, v18)
                for fi36 := 0; fi36 < 12; fi36++ {
                    v3 += v3
                    fmt.Println(fi36, fi36, v3)
                    _ = fi36
                }
                _ = v35
            } else {
                v18 = v18
                var v37 int = v6
                _ = v37
            }
            _ = v34
        } else {
            v3 += (v3 - 2)
            for fi38 := 0; fi38 < 4; fi38++ {
                delete(v5, "foo bar")
                _ = fi38
            }
        }
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    } else {
        for fi39 := 0; fi39 < 15; fi39++ {
            if v4 {
                fmt.Println(v4, v18)
                switch 127 {
                case 32768:
                    switch (^(v6)) {
                    case 127:
                        v3 -= v3
                    case 360:
                        delete(v2, 4294967295)
                    default:
                        fmt.Println(v3)
                    }
                case 9223372036854775807:
                    fmt.Println(v6, fi39, v4)
                case -9223372036854775808:
                    var v40 float32 = 3.14159
                    _ = v40
                }
            } else {
                v3 -= 131
            }
            var v41 bool = v4
            v18 += v18
            _ = fi39
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
