package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 bool; F1 bool; F2 int16 }
func genFn0() int {
    v5 := []int{}
    var v6 string = "a"
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 int, p1 int8) string {
    switch ((("" + "hello") + "") + "a") {
    case "a":
        if (27 != -1) {
            switch genFn0() {
            case 42:
                var v7 int32 = 2147483647
                _ = v7
            case 799:
                if (65535 > 42) {
                    v8 := []int{-1, p0}
                    var v9 uint8 = 255
                    _ = v8
                    _ = v9
                }
            }
            v10 := GenS1{F0: 42}
            _ = v10
        } else {
            fmt.Println(p0, p1)
            if (false && false) {
                for fi11 := 0; fi11 < 7; fi11++ {
                    fmt.Println(p0)
                    for fi12 := 0; fi12 < 12; fi12++ {
                        var v13 uint8 = 255
                        p1 = p1
                        var v14 int = 42
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    var v15 float64 = 1.5
                    _ = fi11
                    _ = v15
                }
                fmt.Println(p0)
                fmt.Println(p1)
            } else {
                if (42 == 1) {
                    fmt.Println(p0, p0)
                } else {
                    var v16 int8 = (p1 + (p1 - (^(p1))))
                    for fi17 := 0; fi17 < 15; fi17++ {
                        for fi18 := 0; fi18 < 12; fi18++ {
                            v19 := []int{-214}
                            p0 -= (p0 * fi18)
                            if true {
                                if (("\n\t" + "foo bar") != ("hello" + "foo bar")) {
                                    var v20 int8 = v16
                                    var v21 float32 = float32(1e10)
                                    v22 := GenS1{F0: 147}
                                    switch ("\n\t" + ("a" + (((("hello" + "") + "a") + "foo bar") + ""))) {
                                    case "":
                                        switch (("\n\t" + "foo bar") + "a") {
                                        case "\n\t":
                                            var v23 int64 = 1
                                            _ = v23
                                        default:
                                            v21 -= v21
                                        }
                                    case "\n\t":
                                        fmt.Println(fi17, v21)
                                    }
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                } else {
                                    v24 := GenS2{F0: 128, F1: -740, F2: 1000}
                                    _ = v24
                                }
                                p1 = (-(p1))
                                var v25 float64 = (((((0.0 + 1.5) / (0.1 / ((-2.25 / (float64((((1.5 + 0.1) * -2.25) / 1.5)) / 1e10)) / 3.14159))) / 0.1) / 1.5) + 0.0)
                                _ = v25
                            } else {
                                p0 = p0
                                _ = len(v19)
                                var v26 int16 = 0
                                _ = v26
                            }
                            _ = fi18
                            _ = v19
                        }
                        p1 += p1
                        _ = fi17
                    }
                    if false {
                        var v27 uint8 = 1
                        var v28 uint32 = 416
                        v29 := GenS3{F0: 2}
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = v16
                }
            }
            fmt.Println(p0, p0)
        }
    default:
        v30 := []int{p0, genFn0(), (p0 * (^(p0))), (p0 & (-(p0)))}
        _ = v30
    }
    if true {
        if (true && false) {
            v31 := []int{genFn0()}
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
            fmt.Println(p0, p0)
            _ = v31
        }
        p0 = p0
        var v32 uint16 = 1
        _ = v32
    }
    p0 += p0
    for fi33 := 0; fi33 < 10; fi33++ {
        for fi34 := 0; fi34 < 5; fi34++ {
            var v35 uint = 18446744073709551615
            var v36 bool = (1000 > -98)
            v37 := GenS2{F0: 255, F1: fi34, F2: 1000}
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        if (!(false && true)) {
            v38 := []int{}
            fmt.Println(p0, p0)
            var v39 int16 = 34
            _ = v38
            _ = v39
        }
        var v40 uint8 = 128
        _ = fi33
        _ = v40
    }
    _ = p0
    _ = p1
    return ("\n\t" + ("" + "foo bar"))
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]string)
    var v3 string = (("" + "") + "")
    v4 := make(map[string]bool)
    for fi41 := 0; fi41 < 1; fi41++ {
        _ = len(v2)
        var v42 bool = (true || false)
        _ = fi41
        _ = v42
    }
    if (!((!(!true)) || true)) {
        v43 := GenS4{F0: (v1 + 42), F1: true}
        _ = v43
    }
    if ((v1 + v1) == 2147483647) {
        v1 -= (v1 >> ((148) & 63))
        fmt.Println(v1, v3)
        v1 += v1
    }
    fmt.Println(v1, v1)
    delete(v4, v3)
    fmt.Println(v1)
    for fi44 := 0; fi44 < 14; fi44++ {
        v45 := GenS5{F0: (genFn0() < fi44), F1: false, F2: 1}
        delete(v2, genFn0())
        for fi46 := 0; fi46 < 4; fi46++ {
            fmt.Println(fi46)
            v1 = v1
            _ = fi46
        }
        _ = fi44
        _ = v45
    }
    fmt.Println(v1)
    for fi47 := 0; fi47 < 2; fi47++ {
        var v48 float64 = 1.5
        var v49 int32 = (v1 | (v1 >> ((356) & 63)))
        v50 := v2[fi47]
        _ = fi47
        _ = v48
        _ = v49
        _ = v50
    }
    switch genFn0() {
    case -1:
        if true {
            v1 = v1
            var v51 int8 = 42
            var v52 int16 = 98
            _ = v51
            _ = v52
        }
    }
    if false {
        for fi53 := 0; fi53 < 2; fi53++ {
            switch "\n\t" {
            case "foo bar":
                for fi54 := 0; fi54 < 17; fi54++ {
                    var v55 uint = 1
                    _ = fi54
                    _ = v55
                }
            default:
                var v56 float32 = 1.5
                _ = v56
            }
            v3 = v3
            var v57 bool = ((127 >= 97) && (3.14159 > 0.1))
            _ = fi53
            _ = v57
        }
        var v58 float32 = (float32((1.5 - (float64(0.1) - (-2.25 + (float64(3.14159) / 3.14159))))) / 1e10)
        _ = v58
    }
    if (true && false) {
        fmt.Println(v1)
        v59 := v2[42]
        _ = v59
    }
    var v60 int8 = 127
    v61 := []int{32767, -367}
    fmt.Println(v3, v3, v60)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v61
}
