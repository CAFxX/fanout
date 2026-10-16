package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 bool = (-27 > -1)
    var v2 float64 = 1e10
    v3 := make(map[string]bool)
    v2 -= v2
    var v4 uint16 = 32767
    var v5 int32 = 0
    v6 := GenS1{F0: 42, F1: ("a" + ("foo bar" + "foo bar"))}
    v7 := GenS1{F0: 18446744073709551615, F1: "foo bar"}
    v4 = (v4 & (v4 * (v4 ^ (v4 * v4))))
    var v8 uint64 = 207
    v1 = (!v1)
    var v9 uint64 = 154
    if true {
        v10 := GenS1{F0: (v8 * 1), F1: ("\n\t" + "a")}
        v2 -= v2
        var v11 int16 = -1
        _ = v10
        _ = v11
    }
    v12 := v3[("hello" + ("a" + (("foo bar" + "foo bar") + ("hello" + ("a" + (("hello" + ((("a" + (("foo bar" + ("foo bar" + "hello")) + ("\n\t" + (((("\n\t" + "\n\t") + "hello") + (("hello" + ("\n\t" + "")) + ("a" + "a"))) + ("" + ("foo bar" + ("hello" + "foo bar"))))))) + "\n\t") + ("hello" + ""))) + "foo bar"))))))]
    v13 := []int{32768, -687, 1, 42}
    v14 := GenS1{F0: v9, F1: ((("hello" + ("" + "\n\t")) + "hello") + ("\n\t" + "a"))}
    if (-2.25 <= (1.5 - float64((float32(0.0) + (3.14159 / 1e10))))) {
        if (42 >= 1) {
            v15 := GenS2{F0: 2147483647, F1: v4}
            for fi16 := 0; fi16 < 10; fi16++ {
                v1 = false
                var v17 bool = false
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 int32 = v5
                fmt.Println(v12, v1, v5)
                _ = fi18
                _ = v19
            }
            _ = v15
        }
    } else {
        for fi20 := 0; fi20 < 20; fi20++ {
            for fi21 := 0; fi21 < 8; fi21++ {
                v22 := GenS2{F0: 2147483647, F1: v4}
                fmt.Println(v1)
                _ = fi21
                _ = v22
            }
            delete(v3, "\n\t")
            if ("a" >= "\n\t") {
                switch ("hello" + "foo bar") {
                case "foo bar":
                    var v23 uint = 42
                    _ = v23
                case "hello":
                    v5 += v5
                case "\n\t":
                    var v24 uint64 = v9
                    _ = v24
                }
            } else {
                if v1 {
                    var v25 bool = true
                    _ = v25
                }
                fmt.Println(v8, v2)
            }
            _ = fi20
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            switch "foo bar" {
            case "a":
                var v27 int64 = 9223372036854775807
                _ = v27
            case "foo bar":
                v12 = (!false)
            case "":
                switch "" {
                case "foo bar":
                    for fi28 := 0; fi28 < 10; fi28++ {
                        fmt.Println(v9, v8, v1)
                        _ = len(v13)
                        v2 += (v2 / ((float64(1.5) / float64(1.5)) / (v2 * -2.25)))
                        _ = fi28
                    }
                }
            }
            var v29 int16 = -642
            _ = fi26
            _ = v29
        }
    }
    v30 := GenS1{F0: 2, F1: ""}
    v9 += (v9 - v9)
    v8 = (v8 << ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
