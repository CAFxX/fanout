package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 string }
func genFn0(p0 int64) int64 {
    switch 128 {
    case 9223372036854775807:
        fmt.Println(p0)
    case -1:
        v4 := []int{32768, -438, -2147483648, -1}
        _ = v4
    case -528:
        if true {
            p0 -= (p0 / ((p0) | 1))
        } else {
            for fi5 := 0; fi5 < 18; fi5++ {
                for fi6 := 0; fi6 < 12; fi6++ {
                    var v7 float32 = 0.0
                    p0 -= p0
                    v8 := GenS1{F0: 65535}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                var v9 int32 = 42
                v10 := []int{(fi5 ^ -1), (fi5 / (((fi5 >> ((0) & 63))) | 1)), fi5}
                _ = fi5
                _ = v10
                _ = v9
            }
            if (!(13 > 128)) {
                v11 := []int{0}
                _ = v11
            }
        }
    }
    var v12 int16 = 32767
    _ = p0
    _ = v12
    return 0
}
func genFn1(p0 int8) uint {
    v13 := GenS1{F0: 0}
    v14 := GenS1{F0: 128}
    v15 := []int{24, -1, 32768}
    if true {
        switch -9223372036854775808 {
        case 42:
            p0 += (-(p0))
        case -90:
            v16 := GenS1{F0: 311}
            _ = v16
        case 1:
            var v17 uint64 = 18446744073709551615
            _ = v17
        default:
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 uint32 = 4294967295
                v20 := GenS2{F0: 745}
                v21 := GenS2{F0: 274}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        var v22 uint32 = 65535
        _ = v22
    }
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return 255
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -9223372036854775808
    var v3 bool = (true && (330 > 355))
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 float32 = (float32((3.14159 * 0.1)) / 3.14159)
        if (4294967295 > 1) {
            v25 := []int{fi23}
            _ = v25
        } else {
            if false {
                v26 := []int{945, 402}
                if false {
                    if v3 {
                        v3 = true
                    } else {
                        v2 = (v2 | (v2 << ((470) & 63)))
                        v3 = (580 > 42)
                        var v27 uint8 = 42
                        _ = v27
                    }
                } else {
                    if len(v26) > 0 {
                        v26[0] = -9223372036854775808
                    }
                    v2 = genFn0(v2)
                    v2 += v2
                }
                if false {
                    var v28 string = ("a" + "hello")
                    var v29 int32 = int32(fi23)
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
            v30 := GenS2{F0: 257}
            _ = v30
        }
        v24 -= v24
        _ = fi23
        _ = v24
    }
    v31 := []int{2147483647, 9223372036854775807, 32767, 32768}
    if true {
        switch (("a" + "hello") + "foo bar") {
        case "\n\t":
            fmt.Println(v2)
        case "foo bar":
            v2 = -416
        default:
            if (145 != 39) {
                v2 -= (v2 - -1)
            } else {
                var v32 uint64 = 938
                _ = v32
            }
        }
        switch "\n\t" {
        case "hello":
            if ((!(((v3 || (982 > 625)) || true) || (40 < 128))) || v3) {
                for fi33 := 0; fi33 < 7; fi33++ {
                    var v34 uint8 = 255
                    v35 := GenS3{F0: 2}
                    delete(v1, fi33)
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 8; fi36++ {
                    fmt.Println(v2, v3)
                    fmt.Println(v3, v2, v3)
                    _ = fi36
                }
            }
        case "a":
            v2 += (^(v2))
        case "":
            v2 = -136
        }
        v2 -= 42
    } else {
        if false {
            v2 = (^(v2))
            for fi37 := 0; fi37 < 11; fi37++ {
                v38 := GenS4{F0: (v2 * genFn0((v2 & 9223372036854775807))), F1: ""}
                v39 := GenS4{F0: (v2 ^ (v2 - (v2 % (((^(v2))) | 1)))), F1: "a"}
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
                _ = fi37
                _ = v38
                _ = v39
            }
            if true {
                var v40 int16 = 1
                var v41 string = ("\n\t" + ("hello" + ""))
                _ = v40
                _ = v41
            } else {
                switch 2147483647 {
                case -1:
                    for fi42 := 0; fi42 < 8; fi42++ {
                        v3 = v3
                        _ = fi42
                    }
                }
                v31 = append(v31, -1)
            }
        }
        v2 = (-(v2))
        v31 = append(v31, 873)
    }
    fmt.Println(v3)
    v2 = v2
    v2 = (v2 * v2)
    _ = len(v1)
    v43 := GenS3{F0: 1}
    var v44 int32 = 68
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v43
    _ = v44
}
