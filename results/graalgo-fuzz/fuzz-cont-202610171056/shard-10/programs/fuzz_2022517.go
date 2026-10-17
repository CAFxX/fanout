package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 uint16, p1 int64, p2 int16) uint16 {
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            switch "hello" {
            case "\n\t":
                var v6 uint8 = 1
                _ = v6
            case "a":
                v7 := []int{fi4, (fi5 * int(p2))}
                _ = v7
            default:
                fmt.Println(p1, p1)
            }
            _ = fi5
        }
        var v8 uint16 = (p0 >> ((0) & 63))
        var v9 int32 = 2147483647
        _ = fi4
        _ = v8
        _ = v9
    }
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    _ = p2
    return p0
}
func genFn1(p0 uint8, p1 bool) int {
    p1 = p1
    var v10 uint32 = 2147483648
    switch 1 {
    case 42:
        fmt.Println(p0)
    case 32768:
        var v11 uint32 = v10
        _ = v11
    case 0:
        v12 := GenS1{F0: (v10 & v10)}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return -1
}

func main() {
    var v1 int = 32767
    var v2 int = (v1 % (((v1 * v1)) | 1))
    var v3 int16 = 42
    if (1 >= 431) {
        fmt.Println(v1, v2, v3)
    } else {
        v13 := []int{(v2 >> ((857) & 63)), genFn1(255, (255 != 255)), v2, v2}
        v3 = (v3 + v3)
        _ = v13
    }
    switch 42 {
    case 9223372036854775807:
        var v14 int16 = v3
        _ = v14
    case 154:
        fmt.Println(v3)
    case 0:
        if false {
            if false {
                v3 = (v3 + (v3 & (-(v3))))
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 uint64 = 4294967295
                    for fi17 := 0; fi17 < 14; fi17++ {
                        v16 -= (v16 >> ((0) & 63))
                        v16 = (v16 & v16)
                        fmt.Println(fi15)
                        _ = fi17
                    }
                    switch "\n\t" {
                    case "hello":
                        for fi18 := 0; fi18 < 12; fi18++ {
                            fmt.Println(v16)
                            fmt.Println(v2, v2, fi15)
                            _ = fi18
                        }
                    case "a":
                        fmt.Println(fi15)
                    case "\n\t":
                        v2 = fi15
                    default:
                        if false {
                            switch 32768 {
                            case 32768:
                                for fi19 := 0; fi19 < 9; fi19++ {
                                    v20 := GenS2{F0: -1, F1: 255}
                                    _ = fi19
                                    _ = v20
                                }
                            default:
                                fmt.Println(fi15, v1)
                            }
                            v3 = (v3 % ((32767) | 1))
                        } else {
                            for fi21 := 0; fi21 < 10; fi21++ {
                                v22 := GenS3{F0: "hello"}
                                _ = fi21
                                _ = v22
                            }
                            var v23 int64 = -1
                            _ = v23
                        }
                    }
                    _ = fi15
                    _ = v16
                }
            }
            fmt.Println(v3)
            v2 = (v2 - 9223372036854775807)
        } else {
            if false {
                if false {
                    v2 = genFn1(128, false)
                    v2 -= v1
                }
                v3 = v3
                var v24 int16 = (v3 + v3)
                _ = v24
            }
        }
    }
    v25 := []int{(v1 + v1)}
    if (42 == 127) {
        var v26 uint64 = 711
        for fi27 := 0; fi27 < 5; fi27++ {
            for fi28 := 0; fi28 < 4; fi28++ {
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                fmt.Println(fi27)
                _ = fi28
            }
            _ = fi27
        }
        _ = v26
    } else {
        var v29 int8 = 100
        _ = v29
    }
    switch (-(v1)) {
    case 91:
        v30 := GenS1{F0: 4294967295}
        _ = v30
    case -1:
        var v31 int = genFn1(0, true)
        _ = v31
    case 347:
        v1 += -126
    }
    if (true && false) {
        v2 -= genFn1(0, (1 < 615))
    }
    v1 = v2
    var v32 int64 = 456
    v3 += 32767
    for fi33 := 0; fi33 < 6; fi33++ {
        v34 := GenS2{F0: 1, F1: 1}
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
}
