package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 string }
type GenS2 struct { F0 uint32; F1 string; F2 string }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int }
func genFn0() bool {
    var v4 string = ("hello" + "foo bar")
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 bool = false
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 42
    v8 := GenS1{F0: 48, F1: 9223372036854775807, F2: (v4 + v4)}
    _ = v4
    _ = v7
    _ = v8
    return (2 >= 65535)
}
func genFn1() uint32 {
    var v9 int8 = 127
    v10 := GenS1{F0: 255, F1: int64(v9), F2: ("hello" + "hello")}
    if (-704 >= 1) {
        v11 := []int{2147483647}
        fmt.Println(v9)
        _ = len(v11)
        _ = v11
    } else {
        fmt.Println(v9)
    }
    v12 := []int{32767}
    _ = v10
    _ = v12
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 1.5
    v3 := []int{-260}
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(fi13, v2)
        fmt.Println(fi13)
        _ = fi13
    }
    var v14 uint16 = 2
    if (false || (-945 < -9223372036854775808)) {
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 float32 = 0.0
            _ = fi15
            _ = v16
        }
        if ((v14 < v14) || true) {
            v17 := GenS2{F0: 102, F1: ("foo bar" + "hello"), F2: "\n\t"}
            v1[("\n\t" + "foo bar")] = true
            v18 := v1["hello"]
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := GenS3{F0: 0, F1: 0}
            var v21 uint8 = 38
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    var v22 int16 = -1
    v23 := GenS3{F0: v14, F1: 127}
    if len(v3) > 0 {
        v3[0] = 0
    }
    if true {
        fmt.Println(v2)
        if (0 == 18446744073709551615) {
            var v24 int32 = -661
            _ = v24
        }
    } else {
        if (true && true) {
            v25 := GenS1{F0: 255, F1: 0, F2: ""}
            _ = v25
        } else {
            switch "foo bar" {
            case "foo bar":
                v2 -= 0.1
            }
        }
        v2 = float32(((-2.25 + 0.0) / 1.0))
        if ("hello" < ("foo bar" + (("\n\t" + "") + ("hello" + "")))) {
            fmt.Println(v2, v14)
            if (!true) {
                fmt.Println(v2)
            } else {
                if true {
                    v1[("a" + ("a" + "foo bar"))] = false
                }
            }
        }
    }
    switch "" {
    case "hello":
        v26 := GenS4{F0: int(v22)}
        _ = v26
    case "":
        if false {
            fmt.Println(v22, v22, v14)
            if false {
                var v27 uint8 = 1
                _ = v27
            } else {
                var v28 uint8 = 115
                fmt.Println(v14)
                _ = v28
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 int32 = int32(fi29)
                var v31 float64 = 0.0
                var v32 uint16 = (v14 >> ((42) & 63))
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        }
    }
    fmt.Println(v22, v22, v2)
    v22 = (v22 >> ((0) & 63))
    v14 = (v14 % (((v14 >> ((281) & 63))) | 1))
    if true {
        v33 := GenS1{F0: 167, F1: 0, F2: ((("\n\t" + "a") + "hello") + "hello")}
        _ = v33
    } else {
        fmt.Println(v22, v14, v2)
        var v34 int8 = -71
        v22 += (v22 & 1)
        _ = v34
    }
    v22 -= 1000
    for fi35 := 0; fi35 < 15; fi35++ {
        delete(v1, "\n\t")
        _ = fi35
    }
    var v36 uint32 = uint32(v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v36
}
