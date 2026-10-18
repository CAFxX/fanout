package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 bool }
func genFn0(p0 uint32, p1 uint8) bool {
    var v7 string = ""
    if (-887 != -1) {
        v8 := []int{}
        _ = v8
    }
    var v9 int = 128
    for fi10 := 0; fi10 < 10; fi10++ {
        if false {
            var v11 uint16 = 0
            _ = v11
        }
        if ((("foo bar" + ("foo bar" + v7)) + v7) >= (("hello" + (v7 + v7)) + ((v7 + (("a" + ("a" + "")) + "a")) + v7))) {
            switch v7 {
            case "":
                switch (-(v9)) {
                case 32768:
                    fmt.Println(p0)
                case 128:
                    var v12 string = v7
                    _ = v12
                case 30:
                    for fi13 := 0; fi13 < 14; fi13++ {
                        var v14 int16 = 855
                        fmt.Println(fi13, p0, p1)
                        _ = fi13
                        _ = v14
                    }
                }
            default:
                switch (v9 & -2147483648) {
                case -1:
                    fmt.Println(fi10, fi10, v7)
                case 9223372036854775807:
                    v9 = v9
                }
            }
        } else {
            switch 669 {
            case 128:
                if false {
                    v15 := GenS1{F0: 580}
                    _ = v15
                } else {
                    var v16 int32 = 42
                    v16 = 2147483647
                    _ = v16
                }
            default:
                v17 := GenS1{F0: 32767}
                _ = v17
            }
            p1 -= p1
            v18 := GenS1{F0: 27}
            _ = v18
        }
        var v19 uint16 = 1
        _ = fi10
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return (!(!(p0 > (p0 << ((2) & 63)))))
}
func genFn1(p0 int, p1 int16, p2 uint64) int32 {
    p1 = (-(p1))
    fmt.Println(p2)
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 bool = (1 > -1)
    v2 := []int{0, -9223372036854775808}
    var v3 int16 = -78
    var v4 uint = 65535
    v5 := make(map[uint32]bool)
    var v6 uint8 = 127
    v1 = ((v6 % (((^(v6))) | 1)) != (v6 | 53))
    var v20 uint32 = 42
    if v1 {
        for fi21 := 0; fi21 < 6; fi21++ {
            if (!v1) {
                var v22 int64 = 42
                switch (fi21 * (fi21 >> ((v4) & 63))) {
                case 4294967295:
                    v4 -= 927
                default:
                    for fi23 := 0; fi23 < 6; fi23++ {
                        v24 := GenS2{F0: v6}
                        _ = fi23
                        _ = v24
                    }
                }
                _ = v22
            } else {
                var v25 bool = v1
                if ("hello" != ("foo bar" + ("hello" + ""))) {
                    v26 := GenS2{F0: v6}
                    fmt.Println(v3, v1)
                    v27 := GenS2{F0: v6}
                    _ = v26
                    _ = v27
                } else {
                    var v28 string = ("hello" + "")
                    var v29 string = (("foo bar" + "hello") + v28)
                    v25 = v1
                    _ = v28
                    _ = v29
                }
                for fi30 := 0; fi30 < 4; fi30++ {
                    _ = len(v2)
                    v3 -= (^(v3))
                    var v31 uint8 = v6
                    _ = fi30
                    _ = v31
                }
                _ = v25
            }
            var v32 int32 = 1
            _ = fi21
            _ = v32
        }
        v6 += 127
    } else {
        v33 := GenS1{F0: 32767}
        _ = v33
    }
    v6 = 128
    if (!(!(-128 != -3))) {
        v34 := v5[v20]
        for fi35 := 0; fi35 < 18; fi35++ {
            v36 := v5[v20]
            v36 = v1
            switch "foo bar" {
            case "\n\t":
                v37 := GenS3{F0: -629, F1: (v3 >> (((v4 & 255)) & 63))}
                _ = v37
            case "foo bar":
                v36 = (-128 == 100)
            default:
                if true {
                    v20 += (v20 << ((v4) & 63))
                } else {
                    v6 -= 147
                    v1 = (true && (!v36))
                }
            }
            _ = fi35
            _ = v36
        }
        v20 += v20
        _ = v34
    } else {
        v38 := GenS4{F0: 643}
        _ = v38
    }
    if v1 {
        for fi39 := 0; fi39 < 18; fi39++ {
            for fi40 := 0; fi40 < 18; fi40++ {
                v1 = false
                v41 := v5[2147483647]
                _ = fi40
                _ = v41
            }
            v4 += (v4 & v4)
            _ = fi39
        }
    } else {
        if v1 {
            var v42 uint16 = 2
            v43 := GenS5{F0: (v3 ^ (v3 / (((v3 & 32767)) | 1))), F1: ((!false) || false)}
            var v44 int64 = -351
            _ = v42
            _ = v43
            _ = v44
        }
        v20 = 506
    }
    var v45 float32 = float32(((float64((-2.25 * float32(1.5))) - 1e10) - 1e10))
    var v46 uint8 = (v6 + (v6 % (((v6 << (((v4 + v4)) & 63))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
