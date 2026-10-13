package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 19
    var v3 int = 128
    v4 := make(map[string]uint32)
    var v5 float32 = (0.0 / (1.5 / (1e10 - -2.25)))
    _ = len(v1)
    v3 -= -557
    v6 := GenS1{F0: 864, F1: 695}
    var v7 int16 = 1
    _ = len(v1)
    var v8 bool = false
    v3 = 42
    v3 = 1
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v9 int64 = -934
        _ = v9
    case "":
        for fi10 := 0; fi10 < 10; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                v3 = fi10
                var v12 bool = (v8 && ((!true) && (4294967295 < 42)))
                var v13 uint32 = 67
                _ = fi11
                _ = v12
                _ = v13
            }
            switch "foo bar" {
            case "foo bar":
                var v14 int64 = -1
                _ = v14
            case "hello":
                var v15 uint16 = 42
                _ = v15
            case "a":
                var v16 int32 = -214
                _ = v16
            default:
                var v17 string = "a"
                _ = v17
            }
            if false {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v19 := GenS2{F0: 423, F1: -28}
                    _ = fi18
                    _ = v19
                }
            } else {
                v4[("hello" + "")] = 2147483648
                var v20 string = ("foo bar" + "hello")
                _ = v20
            }
            _ = fi10
        }
    case "hello":
        v21 := GenS1{F0: 497, F1: 2}
        _ = v21
    }
    v2 += 1
    var v22 int16 = v7
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 int8 = 6
        if v8 {
            var v25 uint16 = 870
            if v8 {
                v7 = v22
                var v26 bool = (((("hello" + ((("a" + ("hello" + ("a" + ("\n\t" + "foo bar")))) + "") + "foo bar")) + "") > "foo bar") || v8)
                v27 := GenS1{F0: 65535, F1: 1}
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 int64 = -1
                _ = fi28
                _ = v29
            }
            _ = v25
        } else {
            if (v3 < -2147483648) {
                v30 := []int{fi23, v3}
                if v8 {
                    var v31 uint8 = v2
                    var v32 int = (fi23 | (v3 - (fi23 % ((fi23) | 1))))
                    _ = v31
                    _ = v32
                } else {
                    var v33 uint8 = 2
                    _ = len(v30)
                    for fi34 := 0; fi34 < 2; fi34++ {
                        v35 := v4[("a" + "a")]
                        _ = fi34
                        _ = v35
                    }
                    _ = v33
                }
                _ = v30
            } else {
                switch v3 {
                case -210:
                    var v36 string = "foo bar"
                    _ = v36
                }
            }
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v8)
    var v37 int = v3
    v37 -= (v37 >> ((2) & 63))
    var v38 bool = (!(v2 < 240))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
