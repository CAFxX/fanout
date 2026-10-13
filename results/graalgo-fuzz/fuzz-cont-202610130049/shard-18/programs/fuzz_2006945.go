package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
func genFn0(p0 uint64, p1 string, p2 int16) uint8 {
    switch (p1 + ("a" + "\n\t")) {
    case "\n\t":
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 255, F1: 2147483647, F2: -1}
            _ = fi7
            _ = v8
        }
    case "":
        v9 := []int{4294967295, 32768}
        _ = v9
    default:
        switch p1 {
        case "\n\t":
            switch 2147483647 {
            case 42:
                v10 := []int{0, 128, 248, 4294967295}
                _ = v10
            }
        case "":
            p1 = p1
        case "hello":
            if false {
                var v11 int8 = 0
                v12 := []int{897, 2147483647, 128}
                p1 = p1
                _ = v11
                _ = v12
            } else {
                p1 = p1
            }
        }
    }
    var v13 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 32767
    var v3 uint8 = 50
    var v4 uint64 = 1
    var v5 int64 = -9223372036854775808
    var v6 uint32 = uint32(v4)
    switch "a" {
    case "a":
        if (1 >= 201) {
            var v14 string = ("\n\t" + ("" + ("" + ("hello" + ("hello" + ("\n\t" + "a"))))))
            _ = v14
        }
    case "":
        v1[-9223372036854775808] = v6
    default:
        v4 -= v4
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := v1[(fi15 + fi16)]
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := GenS2{F0: 0, F1: 1, F2: 127}
    v19 := GenS2{F0: 362, F1: 2147483647, F2: 34}
    if false {
        var v20 bool = false
        var v21 int = 0
        _ = v20
        _ = v21
    }
    switch ((("foo bar" + ("foo bar" + (("\n\t" + "hello") + (("" + ((((("a" + "\n\t") + ("a" + "")) + "foo bar") + "hello") + (("" + "\n\t") + "a"))) + ("" + ("" + "foo bar")))))) + "a") + "hello") {
    case "a":
        var v22 int32 = -525
        _ = v22
    }
    if true {
        switch "hello" {
        case "hello":
            if (v3 >= (^(v3))) {
                var v23 int64 = 857
                for fi24 := 0; fi24 < 11; fi24++ {
                    if true {
                        v23 += -820
                        v5 = 27
                        v25 := []int{(fi24 << ((65535) & 63)), fi24, (-(fi24)), 32768}
                        _ = v25
                    } else {
                        var v26 int = fi24
                        var v27 uint16 = 343
                        v28 := GenS1{F0: 507, F1: 42, F2: 127}
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    v29 := []int{fi24}
                    _ = fi24
                    _ = v29
                }
                for fi30 := 0; fi30 < 19; fi30++ {
                    v31 := []int{9223372036854775807, -2147483648, fi30, fi30}
                    _ = fi30
                    _ = v31
                }
                _ = v23
            } else {
                v3 = (^(v3))
                _ = len(v1)
                v32 := GenS2{F0: 1, F1: -5, F2: 7}
                _ = v32
            }
        default:
            switch -1 {
            case -526:
                var v33 uint = 0
                _ = v33
            case -2147483648:
                v34 := GenS2{F0: 833, F1: 32768, F2: 92}
                _ = v34
            }
        }
        v4 = v4
    }
    switch (((("\n\t" + "\n\t") + "foo bar") + ("foo bar" + "a")) + "\n\t") {
    case "\n\t":
        for fi35 := 0; fi35 < 10; fi35++ {
            var v36 uint8 = genFn0(uint64(v6), "hello", int16(v4))
            _ = fi35
            _ = v36
        }
    default:
        fmt.Println(v2)
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        for fi38 := 0; fi38 < 14; fi38++ {
            fmt.Println(v4, fi38, v3)
            v6 = 355
            _ = fi38
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
