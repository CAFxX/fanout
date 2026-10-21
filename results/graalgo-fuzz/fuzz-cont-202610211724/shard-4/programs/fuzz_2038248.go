package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int16; F1 uint64; F2 int8 }

func main() {
    var v1 bool = (32767 > 42)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    switch "a" {
    case "foo bar":
        v4 := []int{-107, -1, 4294967295}
        _ = v4
    case "\n\t":
        v1 = (true && v1)
    case "a":
        switch ("\n\t" + "a") {
        case "hello":
            fmt.Println(v1)
        case "foo bar":
            if v1 {
                v1 = (!true)
            } else {
                for fi5 := 0; fi5 < 2; fi5++ {
                    var v6 int16 = -71
                    switch (fi5 >> ((999) & 63)) {
                    case 9223372036854775807:
                        for fi7 := 0; fi7 < 9; fi7++ {
                            v8 := GenS1{F0: "a"}
                            _ = fi7
                            _ = v8
                        }
                    case 1:
                        var v9 float64 = ((1e10 + -2.25) * 3.14159)
                        _ = v9
                    }
                    fmt.Println(v1, fi5, v1)
                    _ = fi5
                    _ = v6
                }
            }
        default:
            if true {
                var v10 int64 = 0
                _ = v10
            }
        }
    default:
        var v11 uint64 = 0
        _ = v11
    }
    delete(v2, ("" + "hello"))
    v12 := []int{}
    for fi13 := 0; fi13 < 8; fi13++ {
        v12 = append(v12, 128)
        fmt.Println(v1, v1)
        _ = fi13
    }
    v1 = v1
    fmt.Println(v1)
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint32 = 4294967295
        var v16 uint16 = 624
        v17 := GenS1{F0: "foo bar"}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    if (-2.25 > float64(0.0)) {
        switch "" {
        case "hello":
            _ = len(v2)
        case "\n\t":
            for fi18 := 0; fi18 < 19; fi18++ {
                fmt.Println(v1)
                _ = fi18
            }
        case "foo bar":
            v1 = ("\n\t" > "a")
        }
        v1 = true
    } else {
        if (!(false && (1 != -1))) {
            if false {
                _ = len(v12)
                v19 := GenS1{F0: (((("hello" + "") + "") + "foo bar") + "foo bar")}
                v1 = true
                _ = v19
            }
            v1 = v1
            v20 := GenS2{F0: 1, F1: "hello"}
            _ = v20
        }
        if v1 {
            v1 = (((("" + ("a" + "\n\t")) + ("foo bar" + "foo bar")) + ("a" + "foo bar")) < ("\n\t" + "foo bar"))
        } else {
            if (true || v1) {
                for fi21 := 0; fi21 < 17; fi21++ {
                    switch fi21 {
                    case 128:
                        var v22 uint32 = 185
                        _ = v22
                    }
                    var v23 float64 = 0.1
                    _ = fi21
                    _ = v23
                }
                if true {
                    switch -237 {
                    case 0:
                        fmt.Println(v1)
                    case -9223372036854775808:
                        fmt.Println(v1)
                    }
                } else {
                    var v24 uint = 65535
                    v25 := GenS3{F0: 1, F1: 4294967295, F2: 127}
                    _ = v24
                    _ = v25
                }
            }
            if v1 {
                fmt.Println(v1)
            }
        }
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
