package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }
type GenS3 struct { F0 int64; F1 int32; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint32) uint64 {
    p0 = (p0 % (((p0 ^ (^(p0)))) | 1))
    if ((((float64((0.1 / 1.0)) / (0.1 * 3.14159)) / (1.5 - -2.25)) - float64(3.14159)) == 0.1) {
        fmt.Println(p0)
        v7 := []int{32768, -1, 4294967295}
        _ = v7
    } else {
        v8 := []int{73}
        var v9 int32 = 0
        _ = v8
        _ = v9
    }
    if (false || false) {
        p0 = (p0 >> ((42) & 63))
        v10 := []int{-932, -2147483648}
        _ = v10
    }
    _ = p0
    return 38
}
func genFn1(p0 int8, p1 string, p2 int32) int32 {
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: 255, F1: true}
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 string = ""
            _ = fi13
            _ = v14
        }
        var v15 bool = true
        _ = fi11
        _ = v12
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        if false {
            var v17 int16 = 906
            _ = v17
        } else {
            v18 := []int{fi16}
            var v19 uint64 = 42
            v20 := GenS1{F0: 15, F1: false}
            _ = v18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 int16 = 599
            _ = fi21
            _ = v22
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint = 104
    v2 := []string{"a", ((("foo bar" + ("hello" + "hello")) + (("a" + ("foo bar" + "foo bar")) + "foo bar")) + ""), "\n\t"}
    var v3 string = ""
    var v4 int8 = 127
    var v5 int = 331
    var v6 uint64 = 42
    switch v3 {
    case "a":
        v6 = (v6 * v6)
    case "\n\t":
        v5 = 127
    case "":
        var v23 int16 = -1
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = (v3 + v3)
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 float32 = 0.1
        var v26 int32 = genFn1(v4, v3, 941)
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println(v6, v3, v6)
    fmt.Println(v5, v6)
    fmt.Println(v1, v1)
    if len(v2) > 0 {
        v2[0] = v3
    }
    v27 := GenS1{F0: 128, F1: (true || true)}
    switch (v5 - v5) {
    case 1:
        v4 += (-(v4))
    case 42:
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 14; fi29++ {
                var v30 uint16 = 32767
                fmt.Println(fi28)
                switch (fi28 % ((v5) | 1)) {
                case 724:
                    if (!(!false)) {
                        fmt.Println(fi28, v1)
                        var v31 uint64 = (^(v6))
                        _ = v31
                    }
                }
                _ = fi29
                _ = v30
            }
            switch ((v3 + v3) + ("" + "\n\t")) {
            case "hello":
                for fi32 := 0; fi32 < 5; fi32++ {
                    v5 = (v5 + fi32)
                    if (!(-2.25 == (0.1 / 1.0))) {
                        if ((128 > 1) || ((fi28 <= fi28) || true)) {
                            if (!false) {
                                if true {
                                    switch (fi28 | -2147483648) {
                                    case 127:
                                        v33 := GenS2{F0: 39, F1: (v1 >> (((^(v1))) & 63)), F2: 2147483647}
                                        _ = v33
                                    case -147:
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                    }
                                    v1 += (^(v1))
                                    var v34 bool = (!false)
                                    _ = v34
                                } else {
                                    fmt.Println(v1, fi28)
                                    v6 += (v6 % (((v6 ^ v6)) | 1))
                                    v6 += (v6 + v6)
                                    switch v3 {
                                    case "":
                                        if len(v2) > 0 {
                                            v2[0] = ((v3 + "\n\t") + ("" + v3))
                                        }
                                    case "hello":
                                        fmt.Println(v3)
                                    }
                                }
                                v1 += (v1 & v1)
                                switch fi32 {
                                case 4294967295:
                                    v5 += (v5 & 2147483647)
                                }
                            } else {
                                switch v3 {
                                case "foo bar":
                                    v35 := GenS3{F0: -1, F1: -310, F2: v1}
                                    _ = v35
                                }
                                var v36 string = "\n\t"
                                if len(v2) > 0 {
                                    v2[0] = v36
                                }
                                v5 -= 128
                                _ = v36
                            }
                        }
                    }
                    _ = fi32
                }
            }
            v37 := GenS4{F0: 186}
            _ = fi28
            _ = v37
        }
    case -110:
        var v38 bool = (v3 < "")
        _ = v38
    default:
        v39 := GenS5{F0: (v4 | int8(v1))}
        _ = v39
    }
    var v40 uint = (v1 << ((v1) & 63))
    var v41 uint32 = 871
    v4 += (v4 >> (((v1 << ((uint(v6)) & 63))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
