package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }
type GenS2 struct { F0 int32; F1 int64; F2 int }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint16) uint8 {
    for fi6 := 0; fi6 < 8; fi6++ {
        if ((true || (!(90 >= -128))) || (-2.25 <= 1.5)) {
            v7 := []int{(^(fi6))}
            var v8 string = ("" + "\n\t")
            p0 = (p0 & 208)
            _ = v7
            _ = v8
        } else {
            var v9 uint = 255
            if false {
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 string = (("foo bar" + ("a" + "")) + "a")
                    v9 -= 442
                    _ = fi10
                    _ = v11
                }
                var v12 int32 = 2147483647
                v13 := []int{}
                _ = v12
                _ = v13
            }
            v14 := GenS1{F0: false, F1: "hello", F2: 9}
            _ = v14
            _ = v9
        }
        _ = fi6
    }
    var v15 uint8 = 1
    if true {
        if false {
            var v16 int = 32767
            if true {
                var v17 int8 = 42
                var v18 uint = 920
                _ = v17
                _ = v18
            } else {
                v19 := []int{(v16 & (v16 >> ((112) & 63))), (v16 >> ((255) & 63)), 0, (-(v16))}
                for fi20 := 0; fi20 < 5; fi20++ {
                    if (-822 <= (v16 - int(p0))) {
                        v15 -= 255
                        switch fi20 {
                        case 4294967295:
                            var v21 float32 = float32(float64(0.0))
                            _ = v21
                        case -2147483648:
                            v22 := GenS2{F0: -481, F1: 0, F2: (v16 >> ((119) & 63))}
                            _ = v22
                        default:
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        }
                        v16 -= fi20
                    }
                    for fi23 := 0; fi23 < 5; fi23++ {
                        if ("" == ("" + "")) {
                            var v24 int32 = -130
                            if (true && (1 >= 687)) {
                                if len(v19) > 0 {
                                    v19[0] = v16
                                }
                            }
                            v25 := GenS1{F0: (-776 == 0), F1: "\n\t", F2: 100}
                            v19 = append(v19, 9223372036854775807)
                            v26 := GenS3{F0: v24, F1: p0}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS2{F0: -1, F1: 1, F2: (fi23 | v16)}
                        if true {
                            var v28 int32 = -1
                            _ = v28
                        }
                        _ = fi23
                        _ = v27
                    }
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = fi20
                }
                _ = v19
            }
            v29 := GenS4{F0: 32767, F1: 255}
            _ = v16
            _ = v29
        }
    } else {
        v30 := []int{32767}
        if false {
            v15 -= v15
            var v31 uint8 = (^(v15))
            _ = v31
        } else {
            var v32 int = -471
            var v33 uint32 = 2147483647
            var v34 float32 = float32((3.14159 + 0.0))
            _ = v32
            _ = v33
            _ = v34
        }
        _ = v30
    }
    if false {
        v35 := []int{}
        if (-760 <= -194) {
            p0 = p0
        }
        _ = v35
    } else {
        if ("hello" > ("hello" + (("a" + "foo bar") + "hello"))) {
            if (0.0 > (-2.25 - -2.25)) {
                v15 = 127
                fmt.Println(p0, v15)
                v36 := []int{699, -9223372036854775808}
                _ = v36
            } else {
                var v37 uint16 = (p0 * (^(p0)))
                for fi38 := 0; fi38 < 18; fi38++ {
                    v39 := GenS2{F0: 1, F1: 490, F2: (fi38 - (fi38 >> ((255) & 63)))}
                    fmt.Println(v15, v15)
                    _ = fi38
                    _ = v39
                }
                _ = v37
            }
        }
    }
    p0 += p0
    _ = p0
    _ = v15
    return v15
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    v3 := []uint32{137, 144, 2}
    v4 := make(map[string]int)
    var v5 uint64 = 1
    var v40 uint32 = 42
    if (true || true) {
        fmt.Println(v5, v5)
    } else {
        if len(v2) > 0 {
            v2[0] = ("hello" + "a")
        }
        for fi41 := 0; fi41 < 14; fi41++ {
            for fi42 := 0; fi42 < 4; fi42++ {
                if (((v40 + (v40 - v40)) > v40) && (uint(v40) < 65535)) {
                    _ = len(v4)
                    v40 = v40
                    fmt.Println(v40)
                    v40 = 2147483647
                } else {
                    if (!false) {
                        v40 += (v40 & v40)
                    } else {
                        v43 := GenS5{F0: 380}
                        v44 := GenS3{F0: -1, F1: 1}
                        _ = v43
                        _ = v44
                    }
                }
                _ = fi42
            }
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            _ = fi41
        }
        v45 := v1[(v40 * (v40 & (v40 << ((18446744073709551615) & 63))))]
        _ = v45
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        var v47 float32 = (3.14159 / 0.1)
        _ = fi46
        _ = v47
    }
    if (((3.14159 / (1e10 - 3.14159)) - 1.5) >= (0.1 + float64(1.5))) {
        v5 += v5
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    }
    fmt.Println(v40, v40)
    v48 := GenS3{F0: 42, F1: 42}
    delete(v1, (^(v40)))
    v5 = (v5 + (v5 & (v5 | 2)))
    var v49 uint64 = v5
    v50 := GenS6{F0: -9223372036854775808, F1: 0, F2: v49}
    var v51 int64 = -93
    v40 = (v40 - v40)
    for fi52 := 0; fi52 < 11; fi52++ {
        if ((18446744073709551615 != 65535) && (!true)) {
            v53 := v1[(^(v40))]
            _ = len(v1)
            _ = v53
        } else {
            v1[v40] = fi52
            v49 = (v5 >> ((42) & 63))
            var v54 string = ("" + "hello")
            _ = v54
        }
        v4[((("foo bar" + (("foo bar" + "\n\t") + ("foo bar" + "a"))) + ("\n\t" + "")) + "foo bar")] = fi52
        _ = fi52
    }
    for fi55 := 0; fi55 < 4; fi55++ {
        v56 := GenS7{F0: 65535}
        switch (("a" + "foo bar") + "") {
        case "":
            for fi57 := 0; fi57 < 11; fi57++ {
                v58 := GenS4{F0: int16(fi55), F1: 255}
                _ = fi57
                _ = v58
            }
        default:
            fmt.Println(v40)
        }
        _ = fi55
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
