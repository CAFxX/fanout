package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int64; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint16 }
func genFn0(p0 int8, p1 uint32) int {
    p1 = (p1 << ((1) & 63))
    switch ("hello" + "a") {
    case "foo bar":
        switch (("foo bar" + "\n\t") + ((("\n\t" + "a") + ("\n\t" + (((("foo bar" + "hello") + "a") + ("hello" + "\n\t")) + ("\n\t" + "\n\t")))) + "\n\t")) {
        case "":
            var v5 bool = (-1 == 742)
            _ = v5
        default:
            p1 = p1
        }
    case "a":
        v6 := []int{9223372036854775807, 4294967295, 2147483647}
        _ = v6
    case "\n\t":
        switch -2147483648 {
        case 651:
            var v7 uint16 = 60
            _ = v7
        default:
            v8 := []int{-591, 32768, 805, 664}
            _ = v8
        }
    }
    fmt.Println(p1, p1)
    var v9 int = 4294967295
    v10 := []int{0, 1, (^(v9)), v9}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return v9
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 string = "foo bar"
    var v4 float64 = 3.14159
    if (255 < 558) {
        var v11 int = 9223372036854775807
        if ((v4 / 1.0) > 0.1) {
            fmt.Println(v4)
        }
        _ = v11
    } else {
        v12 := GenS1{F0: 4294967295, F1: 128}
        _ = v12
    }
    v4 -= float64((0.1 + -2.25))
    for fi13 := 0; fi13 < 17; fi13++ {
        if (true || false) {
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 int16 = 0
                _ = fi14
                _ = v15
            }
            v3 = (("" + v3) + v3)
        } else {
            var v16 int = fi13
            _ = v16
        }
        _ = fi13
    }
    if false {
        switch v3 {
        case "\n\t":
            delete(v1, "a")
        case "foo bar":
            v3 = (v3 + v3)
        case "hello":
            if false {
                v3 = (v3 + "hello")
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := GenS1{F0: 65535, F1: 255}
                    var v19 bool = false
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 int32 = 1
                _ = v20
            }
        }
    } else {
        v21 := GenS1{F0: 994, F1: 255}
        v22 := GenS1{F0: 0, F1: 0}
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 string = "a"
        switch v24 {
        case "a":
            v24 = v3
        case "foo bar":
            v25 := GenS2{F0: -9223372036854775808, F1: true}
            _ = v25
        default:
            v26 := []int{fi23}
            _ = v26
        }
        if true {
            for fi27 := 0; fi27 < 8; fi27++ {
                if (v24 == "\n\t") {
                    var v28 uint = 484
                    _ = v28
                }
                v29 := GenS3{F0: 9223372036854775807, F1: 1, F2: 4294967295}
                _ = fi27
                _ = v29
            }
        } else {
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 float32 = 1.5
                v24 = "hello"
                if (!false) {
                    switch v24 {
                    case "a":
                        v32 := GenS4{F0: 637}
                        _ = v32
                    case "hello":
                        var v33 int32 = 0
                        _ = v33
                    default:
                        v34 := []int{genFn0(1, 2147483648), fi30}
                        _ = v34
                    }
                    if (!(!(2 == 0))) {
                        var v35 int16 = 1
                        v36 := []int{(-(fi30)), fi23}
                        _ = v35
                        _ = v36
                    } else {
                        if true {
                            v37 := []int{(^(fi30)), -143, 562, -532}
                            var v38 int16 = 0
                            _ = v37
                            _ = v38
                        } else {
                            var v39 uint = 4294967295
                            fmt.Println(v3, fi23, v4)
                            var v40 float64 = ((float64(float32(v4)) * v4) / 0.1)
                            _ = v39
                            _ = v40
                        }
                        v41 := GenS5{F0: 550, F1: 475}
                        _ = v41
                    }
                    fmt.Println(v31, fi23, fi23)
                }
                _ = fi30
                _ = v31
            }
        }
        _ = fi23
        _ = v24
    }
    if false {
        for fi42 := 0; fi42 < 14; fi42++ {
            if (1.5 <= (1.5 + 0.1)) {
                v43 := GenS2{F0: 1, F1: true}
                v44 := []int{fi42, genFn0(1, 42), 32767, fi42}
                var v45 uint32 = 4294967295
                _ = v43
                _ = v44
                _ = v45
            }
            _ = fi42
        }
        var v46 float64 = float64(0.0)
        v3 = "\n\t"
        _ = v46
    }
    delete(v2, 4294967295)
    v1[((v3 + v3) + v3)] = (1 <= 0)
    var v47 uint32 = 4294967295
    v47 = (v47 & 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
}
