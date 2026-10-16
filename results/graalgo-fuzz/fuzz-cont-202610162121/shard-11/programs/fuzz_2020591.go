package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64) uint8 {
    var v5 uint64 = 18446744073709551615
    var v6 float32 = -2.25
    v6 = v6
    _ = p0
    _ = v5
    _ = v6
    return 107
}
func genFn1() uint16 {
    var v7 string = "foo bar"
    v7 = v7
    switch -2147483648 {
    case -416:
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: true, F1: 100}
                fmt.Println(fi8)
                var v11 int16 = -32768
                _ = fi9
                _ = v10
                _ = v11
            }
            v7 = "hello"
            _ = fi8
        }
    default:
        v12 := []int{127, 4294967295}
        _ = v12
    }
    v7 = "a"
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int16 = 688
        v15 := GenS2{F0: -183, F1: 331, F2: 42}
        var v16 string = ((v7 + ((v7 + ("\n\t" + "a")) + v7)) + v7)
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    _ = v7
    return 528
}

func main() {
    var v1 int16 = 753
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    var v4 string = "foo bar"
    v1 = (v1 - v1)
    var v17 uint64 = 590
    v4 = v4
    fmt.Println(v17)
    v1 = (v1 - (-(v1)))
    if (2147483647 >= 1) {
        var v18 int32 = -2147483648
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 int = (fi19 >> ((255) & 63))
            _ = fi19
            _ = v20
        }
        var v21 uint = 25
        _ = v18
        _ = v21
    } else {
        var v22 float32 = (3.14159 + 1.5)
        _ = v22
    }
    var v23 string = "foo bar"
    for fi24 := 0; fi24 < 4; fi24++ {
        v23 = (v23 + ("hello" + v4))
        var v25 float32 = ((1.5 * 0.0) * 0.1)
        _ = fi24
        _ = v25
    }
    var v26 uint8 = genFn0(704)
    switch "" {
    case "a":
        switch ((v4 + "hello") + "foo bar") {
        case "":
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 int64 = -191
                if (0.1 > 3.14159) {
                    if true {
                        for fi29 := 0; fi29 < 8; fi29++ {
                            var v30 int16 = 1000
                            _ = fi29
                            _ = v30
                        }
                        var v31 bool = false
                        var v32 int16 = (v1 << ((18446744073709551615) & 63))
                        _ = v31
                        _ = v32
                    }
                    if true {
                        if true {
                            var v33 uint32 = 1
                            var v34 bool = true
                            _ = v33
                            _ = v34
                        }
                        var v35 int8 = 70
                        _ = v35
                    } else {
                        v26 = (v26 ^ v26)
                        switch v23 {
                        case "foo bar":
                            fmt.Println(v28, fi27)
                        case "hello":
                            v23 = v23
                        case "\n\t":
                            v36 := GenS3{F0: (v26 & genFn0((^(v28))))}
                            _ = v36
                        default:
                            fmt.Println(v4, v23, v26)
                        }
                    }
                    var v37 int8 = 100
                    _ = v37
                }
                if true {
                    v23 = ("hello" + (v23 + "foo bar"))
                    var v38 int32 = 1
                    _ = v38
                }
                _ = fi27
                _ = v28
            }
        }
    }
    fmt.Println(v1, v17, v4)
    v39 := []int{1, 127, 127, 1}
    var v40 uint32 = 370
    for fi41 := 0; fi41 < 1; fi41++ {
        if (0.0 == 3.14159) {
            v42 := GenS4{F0: 767}
            if ((304 > 1) || (true || (-2147483648 == int32(v17)))) {
                v43 := GenS3{F0: 128}
                _ = v43
            }
            if len(v39) > 0 {
                v39[0] = fi41
            }
            _ = v42
        }
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
        var v44 bool = (uint64(v40) > v17)
        _ = fi41
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
