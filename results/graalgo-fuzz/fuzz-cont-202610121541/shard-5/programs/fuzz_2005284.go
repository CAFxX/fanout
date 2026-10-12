package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint, p1 uint8, p2 uint8) int {
    p2 -= p2
    if (0 != 32767) {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: -55}
            if false {
                var v9 int32 = -2147483648
                p0 = 2
                _ = v9
            }
            var v10 uint8 = p2
            _ = fi7
            _ = v10
            _ = v8
        }
        var v11 uint = 65535
        p1 += (^(p1))
        _ = v11
    } else {
        var v12 int8 = 94
        var v13 uint32 = 2
        if (true || true) {
            switch -9223372036854775808 {
            case -692:
                v14 := []int{-1, 373, -104}
                _ = v14
            case 9223372036854775807:
                var v15 int = int(p1)
                _ = v15
            case -273:
                for fi16 := 0; fi16 < 13; fi16++ {
                    p2 += p1
                    for fi17 := 0; fi17 < 16; fi17++ {
                        var v18 uint64 = 240
                        if (!true) {
                            switch fi16 {
                            case 128:
                                v13 += (v13 % ((42) | 1))
                            case 127:
                                v19 := GenS1{F0: 1}
                                _ = v19
                            }
                            var v20 uint64 = v18
                            if (!(760 != (v13 + v13))) {
                                switch (fi16 & (fi16 & -985)) {
                                case 0:
                                    p1 = 142
                                case 103:
                                    v12 = (v12 - v12)
                                default:
                                    var v21 int16 = -1
                                    _ = v21
                                }
                                var v22 float32 = float32(0.1)
                                _ = v22
                            } else {
                                v12 -= v12
                                v23 := []int{fi16}
                                fmt.Println(v12)
                                _ = v23
                            }
                            _ = v20
                        } else {
                            fmt.Println(fi16, fi16)
                            fmt.Println(fi16, p0)
                        }
                        var v24 int = (fi16 % ((-547) | 1))
                        _ = fi17
                        _ = v18
                        _ = v24
                    }
                    v25 := GenS1{F0: (v12 ^ v12)}
                    _ = fi16
                    _ = v25
                }
            }
        } else {
            v26 := []int{-2147483648, 32767}
            _ = v26
        }
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() int16 {
    v27 := GenS1{F0: 42}
    v28 := []int{128}
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    var v29 int16 = -90
    _ = v27
    _ = v28
    _ = v29
    return v29
}

func main() {
    var v1 uint16 = 42
    var v2 int16 = 1000
    var v3 uint32 = 2147483647
    v4 := make(map[int]int)
    var v5 string = "\n\t"
    v6 := make(map[uint32]int)
    v30 := GenS2{F0: 1}
    if false {
        if (false && true) {
            switch (v5 + v5) {
            case "\n\t":
                v1 = 32767
            }
            var v31 uint8 = 1
            var v32 int = 4294967295
            _ = v31
            _ = v32
        }
    }
    fmt.Println(v5, v3, v1)
    v33 := GenS3{F0: v1, F1: (^(v1))}
    v2 -= 42
    if true {
        fmt.Println(v3, v3, v2)
    } else {
        v1 += 255
        v2 = v2
    }
    var v34 int32 = 690
    var v35 int32 = 2147483647
    for fi36 := 0; fi36 < 10; fi36++ {
        v35 += (v34 / ((v35) | 1))
        _ = fi36
    }
    var v37 int16 = genFn1()
    switch (v5 + v5) {
    case "a":
        v6[(v3 - (v3 | 65535))] = 609
    }
    delete(v4, genFn0(42, 127, 128))
    fmt.Println(v34)
    if true {
        fmt.Println(v1)
    }
    delete(v6, v3)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
