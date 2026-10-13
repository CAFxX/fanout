package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint }
type GenS2 struct { F0 uint64; F1 int32; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 bool; F1 int16; F2 uint32 }
type GenS6 struct { F0 int }
func genFn0(p0 string, p1 uint8) int32 {
    if true {
        if (true && false) {
            for fi6 := 0; fi6 < 1; fi6++ {
                var v7 uint32 = 2147483647
                _ = fi6
                _ = v7
            }
        } else {
            for fi8 := 0; fi8 < 7; fi8++ {
                if ((^(p1)) >= 234) {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        fmt.Println(fi8)
                        _ = fi9
                    }
                    v10 := GenS1{F0: false, F1: 1, F2: 2}
                    _ = v10
                } else {
                    p1 = p1
                }
                _ = fi8
            }
            p1 = p1
            p0 = "foo bar"
        }
        v11 := GenS2{F0: 42, F1: -961, F2: "\n\t"}
        _ = v11
    } else {
        if (!false) {
            v12 := []int{-1}
            var v13 int8 = 127
            for fi14 := 0; fi14 < 10; fi14++ {
                switch int(p1) {
                case -441:
                    switch (("\n\t" + ("foo bar" + ((((p0 + p0) + (p0 + "\n\t")) + p0) + p0))) + "hello") {
                    case "\n\t":
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    case "hello":
                        var v15 uint64 = 4294967295
                        _ = v15
                    case "":
                        v13 += (v13 ^ v13)
                    default:
                        var v16 int32 = -1
                        _ = v16
                    }
                }
                _ = fi14
            }
            _ = v12
            _ = v13
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    return 704
}

func main() {
    var v1 uint16 = 489
    var v2 uint = 0
    v3 := make(map[uint32]int)
    var v4 float32 = 3.14159
    var v5 uint16 = (v1 + (^(v1)))
    var v17 uint64 = 4294967295
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            v17 = v17
            for fi19 := 0; fi19 < 19; fi19++ {
                v4 += 1.5
                _ = fi19
            }
            _ = fi18
        }
        var v20 string = "a"
        if (0 <= (v5 / (((v1 & v5)) | 1))) {
            v21 := []int{}
            var v22 int64 = 9223372036854775807
            if ((0 != 128) || false) {
                _ = len(v21)
                v20 = v20
            }
            _ = v21
            _ = v22
        }
        _ = v20
    }
    if (true && (0 < 2)) {
        var v23 uint = 672
        v24 := GenS3{F0: true}
        _ = v23
        _ = v24
    } else {
        v25 := v3[42]
        v26 := GenS2{F0: v17, F1: 0, F2: "a"}
        for fi27 := 0; fi27 < 20; fi27++ {
            v28 := []int{(-(v25)), v25, -1, v25}
            v29 := v3[2147483647]
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = v25
        _ = v26
    }
    v4 = 0.0
    v30 := []int{}
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := GenS4{F0: 0, F1: true}
        v33 := GenS5{F0: false, F1: -32768, F2: 65535}
        if (-64 >= 1) {
            var v34 float32 = 0.1
            _ = v34
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    v35 := GenS2{F0: 1, F1: -977, F2: (("foo bar" + "") + "a")}
    v36 := v3[65535]
    var v37 float32 = v4
    if true {
        var v38 int8 = 100
        if (("hello" + ("foo bar" + "\n\t")) <= "hello") {
            for fi39 := 0; fi39 < 4; fi39++ {
                var v40 int32 = int32(v17)
                _ = fi39
                _ = v40
            }
            delete(v3, 0)
        }
        _ = v38
    }
    if (!true) {
        v41 := GenS6{F0: v36}
        v5 -= uint16(v2)
        _ = v41
    }
    if true {
        var v42 float32 = 0.1
        v1 += (v1 ^ 491)
        _ = v42
    }
    fmt.Println(v5, v17)
    v37 -= v37
    var v43 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v5
}
