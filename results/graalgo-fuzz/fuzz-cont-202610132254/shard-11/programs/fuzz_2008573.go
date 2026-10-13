package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int8; F1 uint32; F2 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    var v3 int16 = 42
    v3 += 1
    var v4 int64 = 890
    v4 = v4
    v4 += v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v5 uint64 = 42
        if (!((false || false) || true)) {
            v5 = 42
            v6 := v1[0]
            var v7 int = 42
            _ = v6
            _ = v7
        } else {
            v8 := v1[9223372036854775807]
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS1{F0: (v4 - v4), F1: fi9}
                _ = fi9
                _ = v10
            }
            var v11 string = "a"
            _ = v11
            _ = v8
        }
        _ = v5
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        _ = len(v2)
        _ = fi12
    }
    fmt.Println(v4)
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            v4 -= (v4 >> ((2) & 63))
            v3 += v3
            _ = fi14
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint8 = 2
            if true {
                var v17 string = ((("\n\t" + ((("hello" + "") + "foo bar") + ("foo bar" + ((("foo bar" + "\n\t") + "hello") + ("" + ((((("foo bar" + ("hello" + "")) + ((("" + "\n\t") + ("a" + "foo bar")) + "")) + "") + (("\n\t" + ("a" + ("" + ("hello" + ("a" + "hello"))))) + ("foo bar" + "foo bar"))) + ("a" + ("a" + ("hello" + "\n\t"))))))))) + ("a" + "hello")) + "")
                fmt.Println(v17, fi15)
                fmt.Println(fi13, v16, v3)
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        var v18 float32 = 0.0
        _ = fi13
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint16 = 2
        for fi21 := 0; fi21 < 3; fi21++ {
            switch ("hello" + "a") {
            case "":
                v4 += 9223372036854775807
            }
            if (true && (!true)) {
                _ = len(v1)
            }
            if (true || true) {
                if false {
                    fmt.Println(v4, v3, fi19)
                    fmt.Println(v4, fi19, fi19)
                } else {
                    _ = len(v2)
                }
            } else {
                var v22 float64 = 1.5
                v2 = append(v2, (("a" + ("a" + "a")) + "foo bar"))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v22
            }
            _ = fi21
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int16 = 1000
            v3 += (v24 << ((18446744073709551615) & 63))
            v20 += (v20 & v20)
            _ = fi23
            _ = v24
        }
        _ = fi19
        _ = v20
    }
    v3 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch ((("a" + "\n\t") + (("\n\t" + "hello") + ("\n\t" + "a"))) + "") {
    case "a":
        v25 := GenS2{F0: 127, F1: 42, F2: v4}
        _ = v25
    }
    fmt.Println(v3, v4)
    v3 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
