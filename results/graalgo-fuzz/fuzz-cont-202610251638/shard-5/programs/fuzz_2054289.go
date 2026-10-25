package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int64 }
func genFn0(p0 bool, p1 uint16, p2 int32) string {
    p1 = (^(p1))
    var v5 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "a"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("" + "a")
    var v3 uint16 = 2
    var v4 uint64 = 18446744073709551615
    v6 := GenS1{F0: 556, F1: -1, F2: (("\n\t" + "\n\t") + ((v2 + "\n\t") + v2))}
    v7 := GenS1{F0: v3, F1: 100, F2: ((v2 + v2) + "foo bar")}
    switch 0 {
    case -638:
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: 280, F1: 74, F2: v2}
            switch v2 {
            case "foo bar":
                for fi10 := 0; fi10 < 15; fi10++ {
                    switch (fi10 - (fi8 - fi8)) {
                    case 128:
                        v3 = v3
                    }
                    var v11 bool = (796 != 1)
                    _ = fi10
                    _ = v11
                }
            default:
                delete(v1, 783)
            }
            v12 := []int{fi8, fi8, (fi8 | fi8), fi8}
            _ = fi8
            _ = v12
            _ = v9
        }
    case -782:
        fmt.Println(v3, v4)
    case 32767:
        var v13 string = (v2 + (("foo bar" + (("\n\t" + (v2 + v2)) + v2)) + v2))
        _ = v13
    default:
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 int64 = -9223372036854775808
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS1{F0: (v3 ^ (v3 + v3)), F1: 42, F2: ((v2 + v2) + ("a" + ("\n\t" + "hello")))}
                v2 = ("hello" + v2)
                var v18 uint64 = v4
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := GenS2{F0: 255}
            _ = fi14
            _ = v15
            _ = v19
        }
    }
    var v20 string = v2
    v21 := []int{1}
    var v22 int = 32767
    if (467 <= 2) {
        var v23 int16 = -32768
        var v24 string = "hello"
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        _ = len(v21)
        v26 := GenS3{F0: 176, F1: 42}
        for fi27 := 0; fi27 < 2; fi27++ {
            _ = len(v1)
            v28 := GenS1{F0: v3, F1: 63, F2: "foo bar"}
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    if true {
        var v29 int32 = -829
        _ = v29
    } else {
        var v30 int8 = -128
        _ = v30
    }
    fmt.Println(v20, v20, v22)
    if len(v21) > 0 {
        v21[0] = (v22 + v22)
    }
    var v31 uint16 = v3
    v3 += v3
    v3 = 935
    if false {
        v4 += v4
        switch ((("a" + "a") + "\n\t") + ("a" + v2)) {
        case "\n\t":
            switch v20 {
            case "":
                if (float32(0.0) < (1e10 * (0.1 * 0.0))) {
                    switch ("\n\t" + (v20 + "hello")) {
                    case "hello":
                        fmt.Println(v22, v3, v31)
                    }
                    var v32 uint32 = 65535
                    _ = v32
                }
            }
        case "foo bar":
            v2 = v2
        }
    } else {
        v22 = (v22 * 127)
        var v33 int16 = 32767
        _ = v33
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v6
    _ = v7
}
