package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint64; F1 uint; F2 int64 }
func genFn0(p0 uint, p1 int) uint32 {
    v4 := GenS1{F0: 644}
    v5 := GenS2{F0: 0, F1: 695}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 839
}

func main() {
    v1 := []string{"foo bar", "\n\t", ((("a" + "\n\t") + ("" + "\n\t")) + "foo bar")}
    v2 := make(map[int]uint32)
    var v3 int64 = 684
    v6 := GenS1{F0: 42}
    if (-9223372036854775808 <= -9223372036854775808) {
        v7 := GenS2{F0: 65535, F1: 278}
        _ = v7
    }
    var v8 float32 = (float32(1e10) * (((3.14159 + (0.1 * float32(1e10))) - 1.5) * float32(1e10)))
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v9 float64 = (1e10 * 0.0)
    var v10 uint16 = 65535
    var v11 int64 = (^(v3))
    for fi12 := 0; fi12 < 12; fi12++ {
        v11 -= v3
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := GenS1{F0: v10}
            _ = fi13
            _ = v14
        }
        var v15 uint8 = 255
        _ = fi12
        _ = v15
    }
    var v16 int16 = 161
    v8 = 0.0
    var v17 int32 = 2147483647
    for fi18 := 0; fi18 < 20; fi18++ {
        if false {
            v8 += ((-2.25 + 0.0) - (v8 / (v8 - (1e10 / v8))))
            delete(v2, 32768)
            for fi19 := 0; fi19 < 2; fi19++ {
                fmt.Println(v10, fi19, v8)
                v20 := GenS3{F0: (fi18 / ((fi19) | 1))}
                v21 := v2[-348]
                if ((v3 % (((v3 & v11)) | 1)) > v11) {
                    v8 = v8
                    v8 = v8
                } else {
                    fmt.Println(v16, v3, v16)
                    if (128 <= 146) {
                        v1 = append(v1, "hello")
                        switch "hello" {
                        case "":
                            var v22 int = (fi19 % ((int(v11)) | 1))
                            _ = v22
                        case "foo bar":
                            var v23 uint64 = 722
                            _ = v23
                        case "a":
                            v10 += (v10 - (^(v10)))
                        }
                    } else {
                        if false {
                            switch "a" {
                            case "a":
                                v24 := GenS4{F0: (^(v10)), F1: v10}
                                var v25 bool = false
                                var v26 int16 = v16
                                v11 = v11
                                _ = v24
                                _ = v25
                                _ = v26
                            case "hello":
                                if true {
                                    var v27 int32 = -166
                                    v2[(fi19 | fi19)] = (^(v21))
                                    _ = v27
                                } else {
                                    v8 -= float32(1.5)
                                }
                            case "":
                                v28 := GenS3{F0: (-(fi18))}
                                _ = v28
                            default:
                                switch (^(fi19)) {
                                case 480:
                                    _ = len(v1)
                                case 766:
                                    v11 = (v3 | (-(v11)))
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                    v29 := GenS3{F0: fi19}
                                    _ = v29
                                case 1:
                                    v17 += v17
                                default:
                                    var v30 bool = ((!(100 > 97)) || ((v3 + v3) < v11))
                                    _ = v30
                                }
                            }
                        }
                        v16 = 32767
                    }
                    var v31 bool = false
                    _ = v31
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        } else {
            v10 += 0
            switch ("\n\t" + (("hello" + "hello") + "hello")) {
            case "hello":
                fmt.Println(v8, v3)
            }
        }
        v1 = append(v1, "foo bar")
        v1 = append(v1, ("" + "a"))
        _ = fi18
    }
    v32 := GenS5{F0: 2147483647, F1: 255, F2: v11}
    for fi33 := 0; fi33 < 10; fi33++ {
        if ((42 >= -1) && false) {
            v34 := GenS4{F0: (v10 << ((2) & 63)), F1: (^(v10))}
            if (float64(v8) >= v9) {
                v11 = (v3 + int64(fi33))
                v35 := GenS3{F0: (fi33 * 4294967295)}
                _ = v35
            } else {
                var v36 uint32 = genFn0(255, (fi33 % ((fi33) | 1)))
                var v37 uint8 = 214
                _ = v36
                _ = v37
            }
            var v38 string = ("a" + "hello")
            _ = v34
            _ = v38
        }
        var v39 uint64 = 1
        _ = len(v2)
        _ = fi33
        _ = v39
    }
    v40 := GenS6{F0: 722, F1: 2, F2: v3}
    switch ((("\n\t" + "\n\t") + "a") + ("" + "")) {
    case "a":
        var v41 float64 = float64(v8)
        _ = v41
    default:
        v16 -= v16
    }
    switch (("\n\t" + "a") + ("a" + "")) {
    case "\n\t":
        v17 -= v17
    }
    v8 = (v8 * v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v40
    _ = v6
    _ = v8
    _ = v9
}
