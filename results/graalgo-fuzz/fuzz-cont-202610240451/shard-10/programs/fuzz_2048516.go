package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint8 }
func genFn0(p0 int8) string {
    v5 := []int{128, -281, -748}
    p0 += (p0 >> ((2) & 63))
    v6 := GenS1{F0: true, F1: 962}
    switch "a" {
    case "hello":
        if false {
            v7 := GenS1{F0: (!(41 < 127)), F1: 1}
            p0 = (p0 % ((127) | 1))
            _ = v7
        } else {
            for fi8 := 0; fi8 < 9; fi8++ {
                if true {
                    var v9 string = ((("hello" + "a") + "a") + ("foo bar" + ("foo bar" + "a")))
                    var v10 int32 = 1
                    for fi11 := 0; fi11 < 18; fi11++ {
                        v12 := GenS2{F0: 320, F1: 0, F2: 0}
                        if (2 < 32767) {
                            var v13 bool = (127 == 184)
                            fmt.Println(fi11)
                            p0 = -67
                            _ = v13
                        } else {
                            if (!false) {
                                _ = len(v5)
                                v14 := GenS2{F0: 255, F1: 255, F2: 1}
                                _ = v14
                            } else {
                                v10 -= 154
                            }
                            v15 := GenS2{F0: 353, F1: 312, F2: 255}
                            _ = v15
                        }
                        v16 := GenS1{F0: false, F1: v10}
                        _ = fi11
                        _ = v12
                        _ = v16
                    }
                    _ = v10
                    _ = v9
                }
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 uint8 = 31
                    v18 = (v18 >> ((424) & 63))
                    _ = fi17
                    _ = v18
                }
                _ = fi8
            }
        }
    case "\n\t":
        p0 = (p0 >> ((1) & 63))
    case "a":
        p0 -= (p0 / ((-1) | 1))
    default:
        fmt.Println(p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 uint16 = 694
    v4 := make(map[int]int)
    var v19 uint = 1
    v19 -= v19
    fmt.Println(v3, v19)
    switch 0 {
    case -2147483648:
        v20 := GenS1{F0: false, F1: 42}
        _ = v20
    case 32768:
        if true {
            fmt.Println(v19, v19)
            v21 := GenS1{F0: (42 != 4294967295), F1: -2147483648}
            v3 += 42
            _ = v21
        } else {
            switch 127 {
            case -300:
                if true {
                    v19 += 2
                    v22 := []int{0, 1, -818}
                    _ = v22
                }
            case 1:
                v23 := []int{128, 32767}
                _ = v23
            default:
                v19 = 255
            }
            switch "foo bar" {
            case "":
                v3 = 32767
            case "a":
                if true {
                    var v24 float64 = 1e10
                    var v25 int64 = -9223372036854775808
                    for fi26 := 0; fi26 < 4; fi26++ {
                        var v27 bool = true
                        fmt.Println(v27)
                        _ = fi26
                        _ = v27
                    }
                    _ = v24
                    _ = v25
                }
            }
            var v28 int32 = 2147483647
            _ = v28
        }
    }
    v29 := GenS1{F0: (472 < 1), F1: -2147483648}
    fmt.Println(v19)
    var v30 uint16 = (v3 % (((v3 & 255)) | 1))
    v30 = v30
    v30 += v30
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
