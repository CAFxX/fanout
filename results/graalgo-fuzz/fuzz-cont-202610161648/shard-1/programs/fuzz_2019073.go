package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
func genFn0(p0 uint8, p1 uint16) int64 {
    var v7 float32 = 0.1
    var v8 int8 = 96
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 9223372036854775807
}
func genFn1(p0 int, p1 uint64, p2 int64) int {
    p1 += p1
    var v9 int64 = p2
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 36
    v3 := make(map[string]uint32)
    var v4 int16 = 387
    var v5 int8 = v2
    var v6 float64 = 1.5
    var v10 uint64 = 18446744073709551615
    switch 0 {
    case 127:
        var v11 float32 = 1e10
        _ = v11
    case -1:
        var v12 uint16 = 255
        _ = v12
    }
    if (!true) {
        switch (("\n\t" + "a") + "hello") {
        case "":
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := GenS1{F0: 937}
                _ = fi13
                _ = v14
            }
        case "hello":
            if true {
                var v15 uint = 18446744073709551615
                v6 -= float64((3.14159 / 3.14159))
                var v16 int64 = genFn0(127, 255)
                _ = v15
                _ = v16
            } else {
                var v17 int32 = -1
                _ = v17
            }
        case "\n\t":
            v18 := GenS2{F0: "\n\t"}
            _ = v18
        }
    }
    v5 = (v2 | (v2 * -128))
    var v19 string = ("\n\t" + "hello")
    if true {
        var v20 int8 = v5
        _ = v20
    }
    var v21 uint = 2
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := GenS3{F0: (v21 - (v21 + 307))}
        if true {
            fmt.Println(v19)
            if false {
                var v24 int64 = int64(v21)
                var v25 uint32 = 1
                v2 = v5
                _ = v24
                _ = v25
            } else {
                if true {
                    v2 += (v2 % ((v2) | 1))
                    var v26 float32 = float32(v6)
                    _ = v26
                }
            }
        }
        if (9223372036854775807 <= 0) {
            var v27 uint = (v21 % (((v21 - v21)) | 1))
            if (true && true) {
                if (42 == 4294967295) {
                    v4 += (-(v4))
                    _ = len(v3)
                    v28 := GenS2{F0: (v19 + ((v19 + v19) + v19))}
                    _ = v28
                }
                v29 := v3["foo bar"]
                var v30 uint32 = (v29 / (((v29 % ((v29) | 1))) | 1))
                _ = v29
                _ = v30
            } else {
                delete(v3, v19)
            }
            _ = v27
        }
        _ = fi22
        _ = v23
    }
    switch ("foo bar" + v19) {
    case "":
        for fi31 := 0; fi31 < 1; fi31++ {
            var v32 int32 = 0
            var v33 int = (fi31 << ((v21) & 63))
            _ = fi31
            _ = v32
            _ = v33
        }
    case "foo bar":
        for fi34 := 0; fi34 < 2; fi34++ {
            v19 = v19
            _ = fi34
        }
    case "a":
        fmt.Println(v19)
    default:
        var v35 uint8 = 1
        _ = v35
    }
    var v36 uint32 = 42
    var v37 uint = (v21 * (v21 << (((v21 - 564)) & 63)))
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 uint16 = 2
        v39 -= v39
        _ = fi38
        _ = v39
    }
    switch v19 {
    case "hello":
        v6 = 1e10
    default:
        v36 = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
