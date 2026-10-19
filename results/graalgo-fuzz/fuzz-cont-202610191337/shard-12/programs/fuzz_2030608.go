package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int; F1 uint; F2 uint16 }

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 65535
    v3 := []string{}
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    var v6 int8 = 80
    v7 := GenS1{F0: 42, F1: ("foo bar" + "")}
    v2 -= v2
    v8 := GenS2{F0: 0, F1: 1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v2
    fmt.Println(v2, v6, v1)
    if false {
        if false {
            if (v6 == (v6 << ((198) & 63))) {
                var v9 float32 = v1
                v10 := GenS3{F0: 9223372036854775807, F1: 255}
                v11 := GenS2{F0: 1, F1: 1000}
                _ = v10
                _ = v11
                _ = v9
            } else {
                v12 := GenS3{F0: int(v2), F1: 78}
                _ = v12
            }
            var v13 int64 = 0
            v1 -= 1e10
            _ = v13
        }
    }
    if false {
        delete(v4, "a")
    } else {
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 float32 = float32(((3.14159 / float64((float32(0.0) / ((v1 * 1.5) / (0.0 + v1))))) + (3.14159 + 3.14159)))
                var v17 string = ("" + "hello")
                v18 := GenS1{F0: 0, F1: (v17 + v17)}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            var v19 string = ((("hello" + "foo bar") + (("a" + ("\n\t" + (("" + (("" + "\n\t") + "a")) + "foo bar"))) + "a")) + (("" + ("foo bar" + "")) + "hello"))
            _ = fi14
            _ = v19
        }
        v5[32768] = "\n\t"
    }
    fmt.Println(v6, v2)
    var v20 uint16 = 65535
    v21 := v4["\n\t"]
    var v22 uint16 = v20
    var v23 uint = 0
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS4{F0: int(v2), F1: v23, F2: (v22 - 797)}
        switch (((("foo bar" + ("\n\t" + "a")) + ("hello" + "\n\t")) + "hello") + "hello") {
        case "foo bar":
            switch ("a" + (("\n\t" + ((("a" + "\n\t") + ("foo bar" + "\n\t")) + "\n\t")) + "foo bar")) {
            case "a":
                v26 := GenS4{F0: (v21 >> ((4294967295) & 63)), F1: 65535, F2: v20}
                _ = v26
            case "foo bar":
                var v27 int16 = int16(v20)
                _ = v27
            case "hello":
                var v28 int32 = 0
                _ = v28
            }
        }
        _ = fi24
        _ = v25
    }
    v1 = ((0.1 * ((1.5 / v1) / (1.5 / v1))) * v1)
    v29 := GenS4{F0: (v21 & (v21 ^ 128)), F1: v23, F2: v22}
    var v30 string = ("hello" + "\n\t")
    v2 = (v2 & v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
