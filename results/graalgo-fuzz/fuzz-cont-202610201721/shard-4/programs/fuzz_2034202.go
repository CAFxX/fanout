package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32) int64 {
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint16 = 32767
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 2
    if false {
        var v8 int8 = -51
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint32 = 2147483648
            _ = fi9
            _ = v10
        }
        fmt.Println(v8)
        _ = v8
    }
    v11 := GenS1{F0: 4294967295}
    var v12 uint8 = 158
    _ = p0
    _ = v11
    _ = v12
    _ = v7
    return 0
}
func genFn1() uint64 {
    var v13 int32 = -576
    if (!(2 < 905)) {
        v14 := GenS1{F0: 758}
        fmt.Println(v13)
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 10; fi16++ {
                if (!true) {
                    v17 := []int{(fi16 + -1), (fi15 * fi16), (fi15 % ((fi15) | 1)), fi15}
                    _ = v17
                }
                _ = fi16
            }
            _ = fi15
        }
        _ = v14
    } else {
        var v18 int32 = -2147483648
        _ = v18
    }
    v13 = -2147483648
    v13 = v13
    _ = v13
    return 943
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 int64 = -1
    v4 := make(map[uint32]string)
    if (42 > 2) {
        fmt.Println(v3)
    } else {
        delete(v2, 0)
    }
    var v19 float32 = ((float32(1e10) + (1.5 - 0.1)) * (float32((0.1 + 1.5)) + -2.25))
    switch ("a" + "") {
    case "hello":
        for fi20 := 0; fi20 < 2; fi20++ {
            for fi21 := 0; fi21 < 18; fi21++ {
                v22 := GenS1{F0: 4294967295}
                v3 = (-(v3))
                _ = fi21
                _ = v22
            }
            if ((v3 >= 561) || (!true)) {
                var v23 uint = 74
                var v24 bool = false
                _ = v23
                _ = v24
            } else {
                for fi25 := 0; fi25 < 9; fi25++ {
                    v26 := []int{fi20, 2147483647}
                    fmt.Println(v3)
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    _ = fi25
                    _ = v26
                }
                if (true || false) {
                    var v27 uint64 = 839
                    fmt.Println(v3, v3, v27)
                    _ = v27
                }
            }
            var v28 int16 = -32768
            _ = fi20
            _ = v28
        }
    default:
        var v29 uint16 = 2
        _ = v29
    }
    fmt.Println(v19)
    if true {
        v30 := v1[9223372036854775807]
        for fi31 := 0; fi31 < 10; fi31++ {
            _ = len(v1)
            _ = fi31
        }
        for fi32 := 0; fi32 < 10; fi32++ {
            v19 = (0.1 * ((v19 / 3.14159) * v19))
            _ = fi32
        }
        _ = v30
    }
    v33 := GenS2{F0: 207}
    fmt.Println(v19, v3)
    var v34 int8 = 1
    var v35 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
