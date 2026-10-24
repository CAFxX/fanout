package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int8; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS6 struct { F0 int32; F1 string }

func main() {
    var v1 uint32 = 42
    var v2 uint8 = 233
    v3 := make(map[string]bool)
    var v4 uint16 = 32767
    var v5 bool = true
    v6 := v3["foo bar"]
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int64 = 1
        switch -9223372036854775808 {
        case 127:
            var v9 string = ""
            _ = v9
        case -1:
            fmt.Println(v4)
        case 32:
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := GenS1{F0: true}
                v12 := GenS2{F0: 32767}
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            for fi13 := 0; fi13 < 5; fi13++ {
                fmt.Println(v1, v2, v4)
                _ = fi13
            }
        }
        _ = fi7
        _ = v8
    }
    v14 := GenS2{F0: 1}
    switch "foo bar" {
    case "":
        switch 0 {
        case -2147483648:
            if (-128 >= -1) {
                v15 := []int{-49, 128}
                v16 := GenS2{F0: 42}
                _ = v15
                _ = v16
            } else {
                if v5 {
                    v5 = (0 > 32768)
                }
            }
        }
    case "a":
        if (723 >= 4294967295) {
            if (int32(v4) >= -356) {
                v2 -= 0
                var v17 uint64 = 0
                var v18 int64 = -9223372036854775808
                _ = v17
                _ = v18
            }
            v19 := GenS3{F0: 345, F1: 1, F2: (("\n\t" + ("a" + (("foo bar" + "hello") + "foo bar"))) + ("a" + "\n\t"))}
            v20 := GenS3{F0: 738, F1: 81, F2: (("foo bar" + ("hello" + "foo bar")) + ("a" + "\n\t"))}
            _ = v19
            _ = v20
        } else {
            if false {
                v21 := GenS4{F0: (!(v2 >= 2))}
                var v22 uint = 65535
                _ = v21
                _ = v22
            }
        }
    case "foo bar":
        _ = len(v3)
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        if v6 {
            if ((!false) || v5) {
                for fi24 := 0; fi24 < 10; fi24++ {
                    var v25 string = "hello"
                    v26 := GenS5{F0: v4, F1: 42, F2: v1}
                    var v27 uint32 = 4294967295
                    var v28 uint = 65535
                    var v29 uint8 = uint8(v28)
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            }
            switch "hello" {
            case "hello":
                v30 := []int{fi23}
                _ = v30
            default:
                switch (-(fi23)) {
                case 774:
                    var v31 int64 = -9223372036854775808
                    _ = v31
                }
            }
        } else {
            v32 := GenS6{F0: 0, F1: ("\n\t" + "")}
            _ = v32
        }
        v4 = (v4 * v4)
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 int8 = 127
            delete(v3, "foo bar")
            _ = fi33
            _ = v34
        }
        _ = fi23
    }
    delete(v3, "")
    var v35 uint32 = (v1 & (v1 << ((1) & 63)))
    if v6 {
        var v36 float64 = (float64(1e10) / 1.0)
        fmt.Println(v5, v2, v4)
        var v37 int8 = 100
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
