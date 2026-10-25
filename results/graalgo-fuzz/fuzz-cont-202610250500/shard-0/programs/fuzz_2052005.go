package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 0
    v4 := []string{}
    var v5 int16 = v1
    v6 := []string{}
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: fi7, F1: "\n\t", F2: ("a" + "a")}
        var v9 bool = true
        fmt.Println(v9)
        _ = fi7
        _ = v8
        _ = v9
    }
    v1 = (v1 << ((42) & 63))
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint64 = 683
        if false {
            var v12 uint16 = 32767
            _ = v12
        } else {
            v13 := GenS2{F0: 127, F1: 816}
            _ = v13
        }
        if false {
            v14 := GenS2{F0: -1, F1: 65535}
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    switch "\n\t" {
    case "hello":
        v2[v3] = v3
    default:
        if false {
            _ = len(v2)
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 bool = (!(false || true))
                _ = fi15
                _ = v16
            }
        }
    }
    if true {
        for fi17 := 0; fi17 < 3; fi17++ {
            v6 = append(v6, (("foo bar" + ("foo bar" + (("\n\t" + (("\n\t" + ("" + "hello")) + ("" + "foo bar"))) + ("a" + ("hello" + (("\n\t" + "foo bar") + ("hello" + "foo bar"))))))) + "hello"))
            switch (-(fi17)) {
            case 32768:
                var v18 uint8 = 128
                _ = v18
            }
            v19 := GenS2{F0: 0, F1: 65535}
            _ = fi17
            _ = v19
        }
        var v20 int16 = v1
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi21
    }
    var v22 int8 = -128
    v3 = (v3 ^ (v3 * (v3 - v3)))
    for fi23 := 0; fi23 < 20; fi23++ {
        v5 += 1
        var v24 int16 = (v1 * (v1 * 810))
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        for fi26 := 0; fi26 < 1; fi26++ {
            fmt.Println(fi26)
            if false {
                if (!(177 == (v3 % (((v3 ^ 65535)) | 1)))) {
                    v27 := GenS1{F0: (fi25 % ((fi25) | 1)), F1: ((("hello" + "\n\t") + "\n\t") + ""), F2: "\n\t"}
                    fmt.Println(v1)
                    _ = v27
                }
                v5 += v5
                v1 = (v1 | v1)
                fmt.Println(fi25, v22, v1)
                v1 = (v1 * v5)
            }
            _ = fi26
        }
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 uint32 = v3
            v30 := GenS1{F0: (fi25 / ((264) | 1)), F1: "a", F2: ("hello" + "")}
            v29 = (v3 >> ((75) & 63))
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 float64 = 0.0
        _ = fi25
        _ = v31
    }
    switch 627 {
    case 1:
        fmt.Println(v22, v1)
    case -2147483648:
        v32 := GenS2{F0: v22, F1: 0}
        _ = v32
    case 42:
        _ = len(v2)
    }
    var v33 float64 = 0.0
    switch 32767 {
    case 1:
        fmt.Println(v5, v5, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
