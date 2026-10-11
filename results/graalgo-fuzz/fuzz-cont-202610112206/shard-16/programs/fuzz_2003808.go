package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
func genFn0(p0 uint32, p1 int32, p2 int16) int8 {
    var v5 int64 = -549
    switch 32768 {
    case 543:
        var v6 int8 = -50
        _ = v6
    default:
        v7 := []int{}
        _ = v7
    }
    switch ("hello" + "hello") {
    case "foo bar":
        v8 := []int{}
        _ = v8
    case "hello":
        var v9 uint16 = 1
        _ = v9
    }
    var v10 uint8 = 42
    if false {
        var v11 uint = 42
        _ = v11
    } else {
        fmt.Println(p0, v10)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return 42
}
func genFn1() bool {
    var v12 float32 = (((0.1 + 0.1) - (3.14159 * -2.25)) - 0.0)
    v13 := []int{365, 127, 32768}
    v12 = v12
    _ = v12
    _ = v13
    return false
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := []string{"\n\t"}
    v4 := []int{32767, 128}
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint16 = 255
        if false {
            v15 = v15
        } else {
            switch "" {
            case "":
                v16 := GenS1{F0: 9223372036854775807}
                _ = v16
            case "foo bar":
                for fi17 := 0; fi17 < 4; fi17++ {
                    delete(v1, "\n\t")
                    v15 = 42
                    if (((v15 & v15) <= (v15 | v15)) || true) {
                        var v18 int32 = -511
                        var v19 int8 = 45
                        v19 -= (v19 - (v19 >> ((255) & 63)))
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
            default:
                var v20 string = (("" + "a") + (("a" + "a") + "foo bar"))
                _ = v20
            }
            delete(v1, "a")
        }
        _ = fi14
        _ = v15
    }
    v21 := GenS2{F0: 0, F1: 0, F2: -1}
    for fi22 := 0; fi22 < 7; fi22++ {
        switch (fi22 - (fi22 % (((fi22 << ((0) & 63))) | 1))) {
        case 128:
            var v23 uint64 = 42
            _ = v23
        case -9223372036854775808:
            var v24 uint64 = 50
            _ = v24
        default:
            var v25 bool = true
            _ = v25
        }
        _ = fi22
    }
    v26 := GenS2{F0: 42, F1: 128, F2: 42}
    if false {
        var v27 uint64 = 4294967295
        var v28 uint64 = 42
        _ = v27
        _ = v28
    } else {
        var v29 uint16 = 969
        _ = v29
    }
    v30 := GenS1{F0: -1}
    var v31 uint16 = 255
    _ = len(v2)
    var v32 float32 = (0.0 - 3.14159)
    if true {
        if false {
            if true {
                v33 := GenS2{F0: (^(v31)), F1: 42, F2: 2147483647}
                _ = v33
            } else {
                var v34 string = "hello"
                _ = v34
            }
            v32 = v32
        }
        var v35 uint = 4294967295
        if (true && (4294967295 < -189)) {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            var v36 uint8 = 1
            v35 = (v35 << ((v35) & 63))
            _ = v36
        }
        _ = v35
    }
    v31 = (v31 * (^(v31)))
    switch 1 {
    case 0:
        fmt.Println(v32, v31)
    case -9223372036854775808:
        for fi37 := 0; fi37 < 15; fi37++ {
            fmt.Println(v31)
            fmt.Println(v31, v31, fi37)
            _ = fi37
        }
    case -2147483648:
        for fi38 := 0; fi38 < 8; fi38++ {
            for fi39 := 0; fi39 < 5; fi39++ {
                v32 = ((v32 * float32(1.5)) + ((((v32 * v32) - (-2.25 - 1e10)) - v32) / ((v32 / -2.25) - (v32 - (float32(0.1) + float32(1e10))))))
                fmt.Println(v31)
                if false {
                    fmt.Println(fi38, fi38)
                    v40 := GenS3{F0: (fi38 == -859)}
                    var v41 float64 = 1e10
                    _ = v40
                    _ = v41
                }
                _ = fi39
            }
            switch (fi38 ^ fi38) {
            case 1:
                if false {
                    switch -9223372036854775808 {
                    case 598:
                        delete(v1, ("foo bar" + ""))
                    default:
                        switch (((("\n\t" + (("\n\t" + "\n\t") + "hello")) + "hello") + "a") + ("" + ("a" + ("" + "")))) {
                        case "":
                            v32 += (((float32(0.1) - v32) - -2.25) - ((((3.14159 * v32) * 0.1) + v32) * v32))
                        case "hello":
                            v31 -= v31
                        case "foo bar":
                            fmt.Println(v32, v31)
                        default:
                            v42 := GenS4{F0: 4294967295}
                            _ = v42
                        }
                    }
                    var v43 uint32 = uint32(v31)
                    var v44 float32 = v32
                    _ = v43
                    _ = v44
                } else {
                    v45 := GenS3{F0: true}
                    _ = v45
                }
            default:
                v32 = v32
            }
            _ = fi38
        }
    }
    if true {
        v3 = append(v3, "\n\t")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
