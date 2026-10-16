package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int32; F1 uint64; F2 int8 }
func genFn0() bool {
    v4 := []int{0}
    var v5 uint8 = 2
    var v6 int = -2147483648
    for fi7 := 0; fi7 < 18; fi7++ {
        v6 = fi7
        fmt.Println(fi7, v6)
        v8 := GenS1{F0: 42, F1: 42, F2: 127}
        _ = fi7
        _ = v8
    }
    _ = v4
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    v3 := []string{""}
    switch 39 {
    case 127:
        var v9 float64 = -2.25
        _ = v9
    case 1:
        v10 := GenS2{F0: -2147483648}
        _ = v10
    case -580:
        fmt.Println(v1)
    }
    var v11 int = 127
    v1 = v1
    v12 := GenS3{F0: 0}
    if (58 > 39) {
        if (!(!(922 > -32768))) {
            fmt.Println(v1, v11)
        }
        v11 += v11
    } else {
        var v13 bool = v1
        v13 = ((34 < 32767) && ((((4294967295 > (v11 % ((v11) | 1))) && true) && (100 == 1)) && true))
        v11 -= v11
        _ = v13
    }
    fmt.Println(v1, v11)
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 int8 = 0
        _ = fi14
        _ = v15
    }
    var v16 int64 = 1
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS4{F0: (v16 - (-(v16)))}
        _ = fi17
        _ = v18
    }
    v19 := v2[2]
    if v1 {
        for fi20 := 0; fi20 < 12; fi20++ {
            if true {
                switch (("a" + ("" + "foo bar")) + "foo bar") {
                case "":
                    var v21 uint8 = 199
                    _ = v21
                case "foo bar":
                    v19 = true
                }
            } else {
                fmt.Println(v16, v1, v19)
                v11 += (v11 % ((fi20) | 1))
            }
            v19 = (65535 <= 998)
            _ = len(v2)
            _ = fi20
        }
        if true {
            for fi22 := 0; fi22 < 1; fi22++ {
                switch -88 {
                case 32767:
                    fmt.Println(fi22)
                case 144:
                    v11 = fi22
                case 0:
                    if (!((("hello" + "\n\t") + (((((("hello" + "\n\t") + "\n\t") + "") + "\n\t") + "foo bar") + "hello")) >= "foo bar")) {
                        v23 := GenS5{F0: 53, F1: 981, F2: 1}
                        var v24 uint8 = 128
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v23
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 13; fi25++ {
                            v1 = ((3.14159 / 0.1) < float32(((1.5 / -2.25) + 0.1)))
                            if v1 {
                                v26 := GenS3{F0: -122}
                                v3 = append(v3, ("" + ""))
                                var v27 int8 = -1
                                var v28 float32 = float32((-2.25 + float64((((float32(1.5) * 0.0) / -2.25) + (0.1 / 3.14159)))))
                                var v29 string = ""
                                if (1 != 32767) {
                                    fmt.Println(v11)
                                    var v30 uint64 = 0
                                    switch fi25 {
                                    case 2147483647:
                                        v30 = (^(v30))
                                    }
                                    _ = v30
                                } else {
                                    switch (fi22 & int(v27)) {
                                    case -332:
                                        if v1 {
                                            fmt.Println(v11, v11)
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = (v29 + v29)
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = ((v29 + (("foo bar" + "a") + v29)) + v29)
                                            }
                                            switch 42 {
                                            case 4294967295:
                                                v11 += (fi22 + v11)
                                            case 42:
                                                fmt.Println(fi22)
                                            case 127:
                                                v27 = (v27 % (((v27 | (v27 ^ (^(v27))))) | 1))
                                            }
                                            var v31 uint32 = 2147483647
                                            var v32 int8 = v27
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    v11 += -478
                                }
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                var v33 float32 = (0.1 - 0.0)
                                v34 := GenS5{F0: 52, F1: 0, F2: 220}
                                _ = v33
                                _ = v34
                            }
                            _ = fi25
                        }
                        v35 := v2[0]
                        _ = v35
                    }
                }
                _ = fi22
            }
        } else {
            v36 := GenS4{F0: v16}
            v16 -= (v16 >> ((65535) & 63))
            _ = v36
        }
    }
    var v37 float64 = 1e10
    v16 = v16
    var v38 int32 = 951
    fmt.Println(v38, v11, v19)
    fmt.Println(v37, v38, v19)
    switch "foo bar" {
    case "":
        v39 := GenS6{F0: int32(v16), F1: 4294967295, F2: -62}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
}
