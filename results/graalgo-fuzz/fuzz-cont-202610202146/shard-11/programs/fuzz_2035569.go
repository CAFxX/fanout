package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint }
func genFn0() string {
    if true {
        var v7 float32 = 0.1
        v8 := GenS1{F0: 42, F1: 18446744073709551615, F2: 2}
        if (!false) {
            if true {
                fmt.Println(v7)
                var v9 uint16 = 65535
                if true {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        var v11 uint = 602
                        _ = fi10
                        _ = v11
                    }
                } else {
                    v9 += (v9 | 0)
                }
                _ = v9
            } else {
                v7 = v7
                for fi12 := 0; fi12 < 9; fi12++ {
                    v13 := []int{fi12, (fi12 * (^(fi12))), fi12, fi12}
                    if len(v13) > 0 {
                        v13[0] = (fi12 + (fi12 % ((-386) | 1)))
                    }
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = fi12
                    _ = v13
                }
                if true {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        v15 := GenS2{F0: 1000, F1: 661}
                        v16 := []int{32767, 42, 1}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 16; fi17++ {
                        var v18 string = ("" + "hello")
                        _ = fi17
                        _ = v18
                    }
                }
            }
        } else {
            var v19 uint8 = 127
            v20 := GenS3{F0: 557}
            _ = v19
            _ = v20
        }
        _ = v7
        _ = v8
    }
    var v21 bool = (100 != -128)
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 bool = (v21 || false)
        v24 := []int{(fi22 * (fi22 ^ 42))}
        var v25 int8 = 127
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    if v21 {
        switch (("\n\t" + "hello") + ("\n\t" + "a")) {
        case "a":
            v21 = (!(!v21))
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS2{F0: 863, F1: 42}
            v21 = (!(v21 && (0 == 1)))
            var v28 bool = v21
            _ = fi26
            _ = v27
            _ = v28
        }
        fmt.Println(v21)
    }
    _ = v21
    return "\n\t"
}
func genFn1() bool {
    var v29 uint8 = 0
    v29 = v29
    _ = v29
    return false
}

func main() {
    var v1 uint16 = 465
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v4 int32 = -2147483648
    var v5 float64 = (3.14159 / -2.25)
    v6 := []string{"\n\t", "a", "a"}
    fmt.Println(v1, v1, v5)
    v30 := GenS2{F0: 42, F1: 42}
    switch -690 {
    case -452:
        fmt.Println(v5, v1, v5)
    default:
        v5 = (v5 * v5)
    }
    if (((v1 ^ (v1 * 652)) <= v1) && (4294967295 != 42)) {
        var v31 bool = false
        delete(v2, 101)
        _ = v31
    } else {
        switch "foo bar" {
        case "a":
            var v32 uint8 = 2
            _ = v32
        case "foo bar":
            var v33 bool = (2147483647 < v4)
            _ = v33
        default:
            v1 += 65535
        }
        if true {
            for fi34 := 0; fi34 < 2; fi34++ {
                fmt.Println(v4, v5, v5)
                if (!(false || (0 > 4294967295))) {
                    fmt.Println(fi34, v4)
                    var v35 int8 = 105
                    var v36 int = (fi34 % ((fi34) | 1))
                    _ = v35
                    _ = v36
                }
                var v37 string = "hello"
                _ = fi34
                _ = v37
            }
        } else {
            v5 = (v5 / (((1.5 - (v5 + 0.1)) - v5) / ((float64(1e10) * (((v5 + v5) / ((v5 + (float64(-2.25) / float64(1.5))) + (v5 - (v5 - 1.5)))) / v5)) * 1e10)))
            for fi38 := 0; fi38 < 17; fi38++ {
                v1 -= 768
                _ = fi38
            }
        }
        var v39 string = ("hello" + "\n\t")
        _ = v39
    }
    v6 = append(v6, "")
    var v40 uint16 = v1
    v4 -= v4
    v41 := GenS3{F0: 2}
    v1 -= (^(v40))
    for fi42 := 0; fi42 < 20; fi42++ {
        if true {
            var v43 int64 = 0
            switch "hello" {
            case "hello":
                v1 -= (v40 ^ (v1 - (v40 >> ((4294967295) & 63))))
            }
            _ = v43
        }
        for fi44 := 0; fi44 < 11; fi44++ {
            var v45 float32 = (1e10 + float32(0.1))
            if false {
                v5 = 0.1
                var v46 uint32 = 4294967295
                v47 := GenS2{F0: 1000, F1: 18446744073709551615}
                _ = v46
                _ = v47
            } else {
                fmt.Println(v5, fi42)
            }
            v40 = 32767
            if true {
                v48 := GenS2{F0: 1, F1: 465}
                switch fi42 {
                case 4294967295:
                    v5 -= v5
                default:
                    _ = len(v6)
                }
                _ = v48
            } else {
                v4 += 0
            }
            v5 += (v5 / -2.25)
            _ = fi44
            _ = v45
        }
        var v49 int16 = -35
        _ = fi42
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
