package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint, p1 int32) uint32 {
    p1 -= p1
    for fi4 := 0; fi4 < 15; fi4++ {
        fmt.Println(fi4, fi4)
        if (!true) {
            fmt.Println(p0, p1, p0)
        }
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 float64 = float64(((3.14159 + 0.0) * 0.0))
            fmt.Println(fi4, fi5)
            v7 := []int{}
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    v8 := []int{0}
    p1 = (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = v8
    return 65535
}
func genFn1(p0 uint16, p1 uint8, p2 int16) int32 {
    var v9 float32 = 1.5
    v10 := []int{2147483647, 42, int(p0), 32768}
    switch ("foo bar" + "hello") {
    case "foo bar":
        p1 -= (p1 | p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -280
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]uint32)
    var v3 float64 = (1.5 + (3.14159 * (3.14159 * (1.5 * 3.14159))))
    v11 := GenS1{F0: 32768}
    var v12 int16 = 1000
    v13 := v2[1]
    v14 := []int{}
    v12 = v12
    v3 = v3
    for fi15 := 0; fi15 < 12; fi15++ {
        switch int(v13) {
        case 32767:
            if (1 == 1) {
                v1 += v1
                var v16 uint16 = 416
                _ = len(v2)
                _ = v16
            } else {
                if false {
                    v2[(v13 * genFn0(1, -2147483648))] = v13
                    var v17 uint64 = 1
                    _ = v17
                }
            }
        case 1:
            if len(v14) > 0 {
                v14[0] = 727
            }
        }
        v18 := GenS1{F0: 32767}
        var v19 int16 = -1
        _ = fi15
        _ = v18
        _ = v19
    }
    var v20 uint16 = 898
    v3 = (3.14159 - (0.1 / (((v3 / v3) - v3) - v3)))
    var v21 int64 = 9223372036854775807
    v22 := GenS1{F0: 127}
    var v23 string = ("" + ("a" + "hello"))
    fmt.Println(v13, v1)
    if ((v20 - (^(v20))) <= (v20 & (v20 % ((v20) | 1)))) {
        if false {
            v20 -= v20
            var v24 uint16 = (v20 * (v20 + 65535))
            if false {
                v12 = (v12 >> ((762) & 63))
                for fi25 := 0; fi25 < 12; fi25++ {
                    fmt.Println(v23)
                    _ = fi25
                }
            } else {
                fmt.Println(v12, v1, v23)
                for fi26 := 0; fi26 < 7; fi26++ {
                    _ = len(v14)
                    _ = fi26
                }
                delete(v2, v13)
            }
            _ = v24
        }
    }
    switch -1 {
    case 4294967295:
        v27 := GenS1{F0: 32768}
        _ = v27
    case -2147483648:
        if true {
            var v28 uint64 = 1
            for fi29 := 0; fi29 < 18; fi29++ {
                if false {
                    var v30 int8 = -112
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    _ = v30
                } else {
                    var v31 int32 = 2147483647
                    if true {
                        var v32 int32 = int32(v13)
                        for fi33 := 0; fi33 < 18; fi33++ {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            var v34 bool = (v3 != v3)
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    }
                    _ = v31
                }
                _ = fi29
            }
            _ = v28
        }
    }
    fmt.Println(v1, v23, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
