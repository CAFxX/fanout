package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint64 }
type GenS2 struct { F0 uint; F1 int32; F2 bool }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint }
func genFn0(p0 uint, p1 int64) uint32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 int8 = 0
            switch fi8 {
            case -9223372036854775808:
                if (0 == 49) {
                    v9 += 127
                }
            case -2147483648:
                v10 := []int{(fi8 * 1)}
                _ = v10
            case -1:
                var v11 int32 = -1
                _ = v11
            default:
                switch "foo bar" {
                case "a":
                    if false {
                        if true {
                            v9 += (v9 % ((127) | 1))
                        } else {
                            v12 := GenS1{F0: 251, F1: "foo bar", F2: 4294967295}
                            _ = v12
                        }
                        v13 := GenS2{F0: 1, F1: -2147483648, F2: false}
                        _ = v13
                    }
                }
            }
            p0 = p0
            v9 = v9
            var v14 bool = (false && (!(0.1 >= (-2.25 - (0.1 - 3.14159)))))
            var v15 int16 = -472
            p0 = p0
            v15 = -1
            _ = fi8
            _ = v14
            _ = v15
            _ = v9
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(fi16, fi7, fi16)
            var v17 uint16 = 2
            _ = fi16
            _ = v17
        }
        var v18 uint32 = 487
        _ = fi7
        _ = v18
    }
    if (!(!false)) {
        if false {
            for fi19 := 0; fi19 < 17; fi19++ {
                for fi20 := 0; fi20 < 9; fi20++ {
                    if true {
                        p1 = 0
                    } else {
                        p0 += 4294967295
                        v21 := []int{(fi20 | (fi19 - fi19))}
                        _ = v21
                    }
                    _ = fi20
                }
                _ = fi19
            }
            var v22 uint8 = 37
            _ = v22
        } else {
            var v23 int16 = 42
            _ = v23
        }
        v24 := []int{-421, 9223372036854775807}
        p1 = (p1 ^ (p1 % (((p1 % (((p1 + p1)) | 1))) | 1)))
        _ = v24
    }
    if (!true) {
        switch (("\n\t" + ("a" + "hello")) + "foo bar") {
        case "foo bar":
            v25 := GenS1{F0: 128, F1: "\n\t", F2: 0}
            _ = v25
        case "hello":
            if (false && false) {
                if (!true) {
                    if ((p0 << ((p0) & 63)) <= (p0 & 198)) {
                        var v26 string = "a"
                        if false {
                            v26 = "\n\t"
                        }
                        v27 := GenS2{F0: p0, F1: 1, F2: (true || false)}
                        _ = v26
                        _ = v27
                    }
                    p1 = 9223372036854775807
                    v28 := []int{2147483647, int(p0)}
                    _ = v28
                }
            }
        default:
            switch (("" + "a") + ("hello" + (("" + "\n\t") + "a"))) {
            case "foo bar":
                var v29 int = -2147483648
                _ = v29
            }
        }
    }
    p0 -= (^(p0))
    _ = p0
    _ = p1
    return 984
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    var v6 float32 = float32(1e10)
    if (128 <= 2) {
        for fi30 := 0; fi30 < 2; fi30++ {
            if true {
                v1 = (1e10 + 0.0)
                v3[-1] = 33
            }
            _ = fi30
        }
        switch 32768 {
        case -9223372036854775808:
            v31 := GenS2{F0: 291, F1: 2147483647, F2: true}
            _ = v31
        case 123:
            var v32 float64 = (1e10 - (-2.25 / 1.5))
            _ = v32
        }
        fmt.Println(v6, v1)
    } else {
        var v33 float32 = float32(float64(1e10))
        var v34 bool = true
        _ = v33
        _ = v34
    }
    var v35 int = -259
    v36 := GenS1{F0: 2, F1: "hello", F2: 56}
    if (v35 < v35) {
        switch "" {
        case "hello":
            v35 = v35
        case "":
            fmt.Println(v35, v6)
        case "foo bar":
            fmt.Println(v35, v1, v1)
        default:
            if false {
                v6 = v6
                fmt.Println(v35)
            }
        }
    } else {
        switch "a" {
        case "hello":
            v1 = v1
        case "":
            var v37 uint16 = 65535
            _ = v37
        default:
            if (!("a" != "foo bar")) {
                switch "foo bar" {
                case "foo bar":
                    v3[(v35 >> ((448) & 63))] = 42
                case "a":
                    var v38 bool = true
                    _ = v38
                case "hello":
                    v1 = 1e10
                default:
                    v6 = ((v1 * (v1 / (1e10 + float32((1.5 - float64((v1 / float32(-2.25)))))))) + 1e10)
                }
            }
        }
        v39 := GenS3{F0: 2147483648, F1: -527}
        _ = v39
    }
    v40 := GenS4{F0: (1 > 479), F1: 2}
    v41 := GenS5{F0: 255, F1: true}
    var v42 int16 = 0
    var v43 uint8 = 64
    fmt.Println(v35, v43, v6)
    var v44 int8 = 127
    for fi45 := 0; fi45 < 15; fi45++ {
        if false {
            for fi46 := 0; fi46 < 17; fi46++ {
                v47 := []int{v35, v35, (v35 * fi45), fi45}
                if (true || (!false)) {
                    switch "foo bar" {
                    case "a":
                        v42 += (v42 + (v42 | (v42 ^ v42)))
                    }
                    var v48 int16 = v42
                    _ = v48
                } else {
                    if ((-2147483648 >= -451) || true) {
                        v49 := GenS6{F0: 1}
                        var v50 int64 = -914
                        if (false || false) {
                            v51 := GenS4{F0: true, F1: 456}
                            var v52 float32 = 0.0
                            if len(v47) > 0 {
                                v47[0] = (v35 << ((4294967295) & 63))
                            }
                            _ = v51
                            _ = v52
                        }
                        v42 += v42
                        _ = v49
                        _ = v50
                    } else {
                        var v53 int = fi46
                        _ = v53
                    }
                }
                fmt.Println(v35, fi45, v42)
                _ = fi46
                _ = v47
            }
            for fi54 := 0; fi54 < 13; fi54++ {
                fmt.Println(v6, v43, fi54)
                var v55 uint16 = 1
                _ = fi54
                _ = v55
            }
            var v56 int = -1
            _ = v56
        }
        fmt.Println(fi45, v42, v35)
        for fi57 := 0; fi57 < 10; fi57++ {
            delete(v2, v35)
            _ = fi57
        }
        _ = fi45
    }
    v58 := GenS2{F0: 42, F1: 42, F2: false}
    if true {
        v3[(v35 % (((v35 / (((v35 * (v35 / (((^(v35))) | 1)))) | 1))) | 1))] = 211
        var v59 int8 = 1
        _ = v59
    } else {
        if false {
            if (!((v1 / (((float32(-2.25) * 1e10) - 0.0) + v1)) != float32(((3.14159 - (float64((v1 * v6)) / 0.1)) + 0.0)))) {
                v60 := GenS4{F0: ("hello" <= ""), F1: 42}
                _ = v60
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v58
    _ = v6
}
