package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint64, p2 int32) bool {
    fmt.Println(p1, p2)
    var v6 bool = (128 == 99)
    var v7 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (!true)
}

func main() {
    v1 := []uint32{0}
    v2 := make(map[uint32]int)
    var v3 bool = true
    var v4 uint64 = 414
    var v5 int = 42
    v8 := GenS1{F0: "a", F1: 774, F2: 127}
    if ("" == "foo bar") {
        if ((((float32(1.5) / ((0.1 * 0.0) + 1.5)) + 0.0) + 3.14159) != 0.1) {
            v9 := GenS2{F0: (^(v4))}
            v3 = (42 != 473)
            v1 = append(v1, 65535)
            _ = v9
        }
        v3 = false
    } else {
        var v10 uint8 = 42
        switch "foo bar" {
        case "foo bar":
            v4 -= (v4 >> ((997) & 63))
        default:
            switch v5 {
            case 42:
                if v3 {
                    switch (v5 * 4294967295) {
                    case -2147483648:
                        var v11 int64 = 42
                        _ = v11
                    default:
                        for fi12 := 0; fi12 < 13; fi12++ {
                            var v13 float32 = (-2.25 / float32((((-2.25 / (-2.25 + (1.5 / float64(1e10)))) / (0.1 - -2.25)) - ((1e10 + -2.25) + 1e10))))
                            v3 = (!v3)
                            _ = fi12
                            _ = v13
                        }
                    }
                }
            case 127:
                var v14 bool = false
                _ = v14
            case 2147483647:
                v15 := GenS3{F0: 0, F1: -1}
                _ = v15
            }
        }
        _ = v10
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        v3 = v3
        _ = fi16
    }
    var v17 bool = true
    if v17 {
        var v18 int16 = -1
        _ = v18
    }
    if (1 < 49) {
        v19 := GenS1{F0: ("a" + "hello"), F1: 32767, F2: 127}
        v20 := GenS4{F0: -1}
        var v21 uint = 65535
        _ = v19
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 5; fi22++ {
            fmt.Println(v17, v3)
            for fi23 := 0; fi23 < 13; fi23++ {
                v1 = append(v1, 65535)
                switch -715 {
                case 320:
                    fmt.Println(v3, v17, fi22)
                }
                switch ("" + "") {
                case "":
                    if false {
                        var v24 uint16 = 42
                        v4 += (v4 >> ((401) & 63))
                        if v3 {
                            v25 := GenS1{F0: "", F1: v24, F2: -92}
                            if v17 {
                                v2[107] = v5
                            }
                            v26 := GenS5{F0: 2147483647}
                            _ = v25
                            _ = v26
                        }
                        _ = v24
                    }
                }
                _ = fi23
            }
            _ = fi22
        }
    }
    switch ("\n\t" + "hello") {
    case "a":
        if (169 != 42) {
            var v27 float32 = 1.5
            _ = v27
        } else {
            _ = len(v1)
            switch ("\n\t" + ((("foo bar" + (("hello" + "\n\t") + "foo bar")) + "foo bar") + "a")) {
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "hello":
                var v28 uint64 = (^(v4))
                _ = v28
            case "\n\t":
                var v29 bool = v17
                _ = v29
            default:
                var v30 string = (("hello" + "foo bar") + "")
                _ = v30
            }
        }
    case "\n\t":
        v31 := GenS3{F0: -1, F1: 0}
        _ = v31
    case "hello":
        var v32 int16 = 42
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        delete(v2, 4294967295)
        var v34 uint16 = 0
        _ = fi33
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
