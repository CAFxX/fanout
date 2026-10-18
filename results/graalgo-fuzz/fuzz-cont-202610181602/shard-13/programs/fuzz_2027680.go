package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
type GenS3 struct { F0 uint8; F1 string; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0() uint16 {
    var v6 int64 = -1
    v7 := GenS1{F0: (("foo bar" + "a") + "a")}
    if false {
        v6 -= v6
    }
    v8 := GenS2{F0: 128, F1: "foo bar", F2: -2147483648}
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[string]string)
    var v3 uint16 = 1
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    v1 = (v1 >> ((0) & 63))
    fmt.Println(v3)
    fmt.Println(v3, v3)
    fmt.Println(v3)
    v9 := GenS3{F0: 128, F1: ("hello" + "\n\t"), F2: 42}
    v5[42] = "\n\t"
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int32 = 0
        v11 = -557
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := []int{fi12, -1}
        v14 := GenS4{F0: 32767, F1: 2}
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 = v3
    var v15 uint16 = (v3 << ((2) & 63))
    for fi16 := 0; fi16 < 5; fi16++ {
        if true {
            v3 = (^(v3))
            v1 += -756
            v1 += 0
        }
        _ = fi16
    }
    var v17 uint16 = v3
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            v17 = v17
            if (-1 <= -962) {
                var v20 uint8 = 18
                v20 -= (v20 + 1)
                _ = v20
            } else {
                v3 = 65535
                v21 := []int{fi18}
                var v22 int16 = -32768
                switch (((((("a" + "a") + (("" + ((((("hello" + "\n\t") + "foo bar") + "\n\t") + "hello") + "hello")) + "\n\t")) + "\n\t") + "a") + "\n\t") + "\n\t") {
                case "a":
                    if (!(-128 < 127)) {
                        v22 = -1
                        v15 -= v3
                    } else {
                        v3 = v3
                        switch v1 {
                        case 127:
                            v17 = 442
                        default:
                            if false {
                                switch (("hello" + "hello") + "a") {
                                case "":
                                    switch "\n\t" {
                                    case "hello":
                                        var v23 float64 = ((3.14159 * 0.0) + 1.5)
                                        _ = v23
                                    case "\n\t":
                                        var v24 uint = 736
                                        _ = v24
                                    case "a":
                                        var v25 int16 = (^(v22))
                                        _ = v25
                                    default:
                                        var v26 uint = 752
                                        _ = v26
                                    }
                                }
                            }
                        }
                        var v27 float64 = -2.25
                        _ = v27
                    }
                default:
                    v17 = (v17 - (v3 & (v17 / (((^(v3))) | 1))))
                }
                _ = v21
                _ = v22
            }
            var v28 int64 = -792
            _ = fi19
            _ = v28
        }
        var v29 int8 = -6
        _ = fi18
        _ = v29
    }
    var v30 bool = (-1 < 0)
    switch (("" + "\n\t") + (("foo bar" + "") + ("\n\t" + "foo bar"))) {
    case "foo bar":
        var v31 uint16 = (v17 & (v15 | (v17 * 42)))
        _ = v31
    case "\n\t":
        var v32 uint32 = 65535
        _ = v32
    case "hello":
        v33 := []int{v1, v1, v1}
        _ = v33
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        delete(v5, 0)
        if ("\n\t" < "foo bar") {
            if (v17 == (v15 & v17)) {
                for fi35 := 0; fi35 < 4; fi35++ {
                    switch ("" + "") {
                    case "hello":
                        var v36 float32 = ((((0.0 / float32((0.0 + (1.5 - 3.14159)))) - (3.14159 + 0.1)) / 1.0) * -2.25)
                        _ = v36
                    }
                    v37 := []int{(v1 | (^(v1))), (fi34 + fi35)}
                    switch ("foo bar" + ("hello" + "")) {
                    case "hello":
                        if ((v17 % ((v3) | 1)) == v3) {
                            v3 -= v15
                        } else {
                            delete(v5, 4294967295)
                            var v38 uint8 = 255
                            _ = v38
                        }
                    }
                    _ = fi35
                    _ = v37
                }
            }
        } else {
            var v39 bool = (false || v30)
            var v40 uint8 = 2
            _ = v39
            _ = v40
        }
        var v41 float64 = (3.14159 * 1e10)
        _ = fi34
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
