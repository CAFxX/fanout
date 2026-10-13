package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint8 }
func genFn0() int16 {
    var v7 int16 = 1000
    switch "\n\t" {
    case "":
        var v8 bool = true
        _ = v8
    case "foo bar":
        v7 = -372
    case "hello":
        var v9 int16 = (v7 >> ((394) & 63))
        _ = v9
    }
    v10 := GenS1{F0: 1}
    v11 := GenS2{F0: 18446744073709551615, F1: 145, F2: 59}
    _ = v10
    _ = v11
    _ = v7
    return (v7 + (v7 + 1))
}
func genFn1() string {
    var v12 uint64 = 18446744073709551615
    for fi13 := 0; fi13 < 1; fi13++ {
        v12 = v12
        fmt.Println(fi13, fi13)
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 bool = (1 < -32768)
            v12 -= 42
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    _ = v12
    return ("\n\t" + "foo bar")
}

func main() {
    var v1 int8 = 100
    var v2 uint32 = 637
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    var v5 bool = (false && (1e10 < 1.5))
    v6 := make(map[string]int)
    v5 = (!v5)
    for fi16 := 0; fi16 < 16; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            v1 = (v1 >> ((18446744073709551615) & 63))
            var v18 uint32 = 2147483647
            v1 -= v1
            _ = fi17
            _ = v18
        }
        if v5 {
            v19 := []int{int(v1), (fi16 + 32768), fi16}
            var v20 int16 = 1
            _ = v19
            _ = v20
        } else {
            v6[("foo bar" + ((("hello" + "foo bar") + "") + "\n\t"))] = fi16
        }
        _ = len(v3)
        _ = fi16
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v1 -= v1
        v1 = 1
        _ = fi21
    }
    fmt.Println(v1)
    if false {
        v22 := GenS1{F0: 712}
        v1 += v1
        _ = v22
    } else {
        for fi23 := 0; fi23 < 13; fi23++ {
            switch (^(fi23)) {
            case 32767:
                v1 = v1
            default:
                var v24 int64 = -9223372036854775808
                _ = v24
            }
            var v25 int16 = 350
            var v26 int = fi23
            _ = fi23
            _ = v25
            _ = v26
        }
        switch "" {
        case "":
            v27 := []int{int(v2), 32767, 32767}
            _ = v27
        default:
            var v28 int16 = 42
            _ = v28
        }
        var v29 int8 = 78
        _ = v29
    }
    var v30 int32 = -387
    var v31 uint16 = 1
    var v32 float32 = 1.5
    var v33 int64 = 841
    v5 = (true || true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
