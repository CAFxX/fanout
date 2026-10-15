package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 bool; F1 int8; F2 uint32 }
func genFn0() uint16 {
    if (42 < -336) {
        for fi4 := 0; fi4 < 9; fi4++ {
            if true {
                v5 := GenS1{F0: 64, F1: "", F2: 42}
                _ = v5
            }
            var v6 bool = false
            v6 = true
            _ = fi4
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        if true {
            if true {
                v8 := GenS1{F0: -2147483648, F1: (("foo bar" + ((("" + "hello") + ("foo bar" + "a")) + (("hello" + ("foo bar" + "a")) + "\n\t"))) + ""), F2: 1}
                var v9 uint64 = 0
                _ = v8
                _ = v9
            }
            var v10 uint32 = 2
            _ = v10
        } else {
            for fi11 := 0; fi11 < 15; fi11++ {
                switch (fi11 / ((fi11) | 1)) {
                case 0:
                    var v12 float32 = 0.1
                    _ = v12
                case 32767:
                    var v13 int32 = 42
                    _ = v13
                case -9223372036854775808:
                    if false {
                        var v14 uint8 = 1
                        var v15 bool = (false || (false || (255 != 32767)))
                        v16 := GenS1{F0: 1, F1: "a", F2: 42}
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
                _ = fi11
            }
            var v17 int16 = 32767
            _ = v17
        }
        _ = fi7
    }
    var v18 int = 1
    _ = v18
    return 1
}

func main() {
    v1 := []uint32{0, 0, 720}
    var v2 int32 = 42
    var v3 int = 4294967295
    fmt.Println(v3, v2)
    fmt.Println(v3, v2)
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 uint16 = 855
        v2 = (v2 * v2)
        var v21 uint64 = 1
        _ = fi19
        _ = v20
        _ = v21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 int32 = (-(v2))
    v23 := GenS2{F0: true, F1: 11, F2: 42}
    if false {
        var v24 float64 = 0.0
        v25 := GenS1{F0: (v22 ^ (v22 << ((65535) & 63))), F1: ("a" + ("hello" + "")), F2: 1}
        _ = v24
        _ = v25
    }
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    if true {
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 uint = 169
            _ = fi26
            _ = v27
        }
        var v28 bool = (4294967295 > 241)
        _ = v28
    } else {
        var v29 int16 = 1
        var v30 string = "foo bar"
        fmt.Println(v3)
        _ = v29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
