package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int) uint32 {
    v7 := GenS1{F0: "a"}
    var v8 uint = 65535
    v9 := GenS1{F0: ("\n\t" + ("" + "a"))}
    p0 -= 127
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint32, p1 uint64, p2 int64) string {
    v10 := GenS1{F0: "a"}
    fmt.Println(p0, p2)
    var v11 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return ("a" + "hello")
}

func main() {
    var v1 uint16 = 65535
    var v2 bool = (-32768 != -32768)
    var v3 int8 = 1
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]uint32)
    var v6 string = "\n\t"
    if (208 < 42) {
        fmt.Println(v1)
        v6 = v6
    } else {
        for fi12 := 0; fi12 < 16; fi12++ {
            if true {
                for fi13 := 0; fi13 < 1; fi13++ {
                    v14 := GenS2{F0: (v1 & 65535), F1: 18446744073709551615}
                    var v15 int8 = -1
                    _ = len(v4)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            } else {
                v1 = 0
            }
            _ = fi12
        }
        v1 = (v1 + v1)
        v6 = ((((v6 + "foo bar") + "\n\t") + v6) + v6)
    }
    if false {
        var v16 int = -126
        fmt.Println(v6)
        v16 = (v16 >> ((42) & 63))
        _ = v16
    } else {
        switch -23 {
        case -400:
            for fi17 := 0; fi17 < 20; fi17++ {
                delete(v5, 2147483648)
                var v18 uint32 = genFn0(2147483647)
                _ = fi17
                _ = v18
            }
        case -156:
            v1 = (^(v1))
        default:
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 float64 = 1e10
                v21 := []int{(fi19 << ((18446744073709551615) & 63)), 2147483647}
                _ = fi19
                _ = v20
                _ = v21
            }
        }
    }
    var v22 bool = ((2 <= (v1 % ((692) | 1))) || (812 == 996))
    if v2 {
        v2 = v2
        if false {
            fmt.Println(v3, v2, v1)
            var v23 bool = (v2 || v22)
            for fi24 := 0; fi24 < 1; fi24++ {
                if false {
                    v25 := []int{(fi24 ^ 1), (fi24 % ((fi24) | 1))}
                    for fi26 := 0; fi26 < 1; fi26++ {
                        switch (fi26 - fi26) {
                        case 0:
                            if (171 <= v1) {
                                v2 = v22
                                v1 = 42
                            }
                        }
                        _ = fi26
                    }
                    _ = v25
                }
                _ = fi24
            }
            _ = v23
        } else {
            var v27 int = 678
            v27 = v27
            v28 := []int{0}
            _ = v27
            _ = v28
        }
        v29 := GenS2{F0: (^(v1)), F1: 18446744073709551615}
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v3 -= (v3 << ((4294967295) & 63))
        _ = fi30
    }
    fmt.Println(v3, v3, v2)
    v22 = v22
    switch (v6 + v6) {
    case "hello":
        var v31 int = -2147483648
        _ = v31
    }
    v32 := v4[2147483648]
    var v33 float32 = 0.1
    fmt.Println(v6)
    fmt.Println(v32, v33)
    v33 += v33
    var v34 uint = 255
    for fi35 := 0; fi35 < 5; fi35++ {
        v22 = (true || (float32(1e10) < (v33 + (0.0 / (v33 / 3.14159)))))
        _ = fi35
    }
    switch v6 {
    case "":
        v1 -= 0
    case "a":
        var v36 uint64 = 1
        _ = v36
    case "hello":
        delete(v4, v32)
    }
    fmt.Println(v3, v22, v33)
    if v2 {
        v32 = (v32 | (v32 | v32))
        for fi37 := 0; fi37 < 9; fi37++ {
            v2 = true
            var v38 float64 = (1e10 / (-2.25 + 0.0))
            v39 := GenS3{F0: 1}
            _ = fi37
            _ = v38
            _ = v39
        }
        for fi40 := 0; fi40 < 14; fi40++ {
            var v41 int8 = -109
            for fi42 := 0; fi42 < 8; fi42++ {
                var v43 uint64 = 42
                var v44 int8 = 42
                _ = len(v4)
                v45 := []int{fi42, fi42}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            v3 = (v3 + v41)
            _ = fi40
            _ = v41
        }
    } else {
        v46 := v5[v32]
        for fi47 := 0; fi47 < 9; fi47++ {
            if (v1 == v1) {
                v6 = (("a" + v6) + v6)
                fmt.Println(v3)
                fmt.Println(v6, v34, fi47)
            }
            _ = fi47
        }
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
