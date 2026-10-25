package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 int8 }

func main() {
    v1 := []string{"a", ("foo bar" + "hello")}
    var v2 float64 = 0.1
    v3 := make(map[int]string)
    var v4 uint32 = 1
    v5 := make(map[string]bool)
    var v6 uint16 = 65535
    var v7 bool = true
    _ = len(v3)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    if false {
        v4 -= (v4 ^ v4)
        var v8 float64 = ((v2 + 0.1) - v2)
        _ = v8
    }
    if false {
        v2 = (v2 / v2)
        v6 = (v6 ^ v6)
        v6 += v6
    } else {
        var v9 uint32 = v4
        v9 -= v4
        _ = v9
    }
    if ((("foo bar" + (((("\n\t" + (("\n\t" + ("hello" + "hello")) + (("\n\t" + "") + ("a" + (("" + "\n\t") + ((("" + "foo bar") + "hello") + ("\n\t" + "\n\t"))))))) + "") + "hello") + "foo bar")) != "") && ((v4 >> ((579) & 63)) != (v4 + v4))) {
        var v10 int32 = 2147483647
        v11 := GenS1{F0: -8, F1: -520}
        v5["hello"] = (!(("foo bar" + "foo bar") == ("a" + "\n\t")))
        _ = v10
        _ = v11
    }
    if (852 >= 1) {
        var v12 float64 = float64(1.5)
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint16 = v6
        for fi15 := 0; fi15 < 5; fi15++ {
            switch (fi13 ^ int(v6)) {
            case 9223372036854775807:
                fmt.Println(v2, fi13, v6)
            case 2147483647:
                var v16 int16 = 32767
                _ = v16
            }
            _ = fi15
        }
        switch "foo bar" {
        case "foo bar":
            v17 := GenS1{F0: 0, F1: (fi13 - fi13)}
            _ = v17
        case "":
            v18 := GenS2{F0: 32767, F1: "a", F2: 0}
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint64 = 42
        v7 = true
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
