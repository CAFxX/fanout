package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int64 }
func genFn0(p0 int32, p1 int32) uint8 {
    v4 := []int{}
    if ((!(!("a" != (("" + "a") + "")))) && false) {
        for fi5 := 0; fi5 < 7; fi5++ {
            if false {
                p1 = p1
            } else {
                var v6 float32 = -2.25
                _ = v6
            }
            p0 = (p1 - p0)
            _ = fi5
        }
        var v7 bool = false
        var v8 uint16 = 1
        _ = v7
        _ = v8
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = p1
    _ = v4
    return 127
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[uint32]string)
    v3 := []string{"", "foo bar"}
    var v9 int32 = 42
    var v10 int32 = (v9 | -1)
    var v11 uint8 = genFn0(0, 42)
    switch ("foo bar" + ("hello" + "a")) {
    case "":
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint32 = 95
            var v14 int8 = -2
            v2[v13] = ("" + "hello")
            _ = fi12
            _ = v13
            _ = v14
        }
    case "foo bar":
        var v15 int32 = 141
        _ = v15
    }
    v9 = v10
    var v16 uint8 = (v11 << ((42) & 63))
    for fi17 := 0; fi17 < 12; fi17++ {
        if (65535 > 1) {
            var v18 uint16 = 1
            switch "foo bar" {
            case "a":
                v19 := GenS1{F0: -32768, F1: false, F2: 734}
                _ = v19
            }
            _ = v18
        }
        switch 4294967295 {
        case 637:
            v11 = v16
        }
        _ = fi17
    }
    v20 := GenS1{F0: 42, F1: false, F2: 42}
    if len(v3) > 0 {
        v3[0] = ("foo bar" + "")
    }
    var v21 int32 = v10
    v10 = -2147483648
    fmt.Println(v9, v21)
    switch 2147483647 {
    case 127:
        if len(v1) > 0 {
            v1[0] = ("" + "\n\t")
        }
    }
    var v22 int8 = 100
    v9 = -1
    for fi23 := 0; fi23 < 3; fi23++ {
        v22 += v22
        _ = fi23
    }
    if (false && (!false)) {
        v11 = v11
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
