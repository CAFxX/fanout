package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 int; F1 string; F2 int16 }
func genFn0(p0 int64) int8 {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 918}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        if ((!(128 != 2)) || (!(("hello" + "") < "foo bar"))) {
            if (48 < 1) {
                p0 -= 0
                var v8 int = fi7
                v9 := GenS2{F0: (1 >= 720), F1: 0, F2: "hello"}
                _ = v8
                _ = v9
            }
        }
        var v10 int8 = 110
        var v11 int32 = -1
        _ = fi7
        _ = v10
        _ = v11
    }
    v12 := []int{2147483647, 128}
    var v13 int8 = -68
    var v14 uint32 = 287
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return v13
}

func main() {
    v1 := []uint32{293, 42, 42}
    var v2 int32 = -1
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    switch ("hello" + (("" + "a") + "foo bar")) {
    case "hello":
        v2 = -623
    case "a":
        for fi15 := 0; fi15 < 9; fi15++ {
            switch "hello" {
            case "":
                for fi16 := 0; fi16 < 2; fi16++ {
                    _ = len(v4)
                    var v17 string = (("\n\t" + ("a" + "\n\t")) + "hello")
                    v17 = (v17 + ("hello" + v17))
                    v18 := GenS2{F0: ((1e10 - (1.5 - (0.1 + 3.14159))) == (0.1 * 0.1)), F1: 892, F2: "\n\t"}
                    switch v17 {
                    case "foo bar":
                        v17 = "\n\t"
                    case "\n\t":
                        v17 = "foo bar"
                    case "":
                        if false {
                            var v19 float64 = -2.25
                            v20 := v3[19]
                            _ = v19
                            _ = v20
                        }
                    }
                    if false {
                        var v21 int64 = -791
                        var v22 int64 = v21
                        v17 = ("hello" + ("hello" + (v17 + v17)))
                        delete(v3, 2)
                        v23 := GenS3{F0: (fi16 | -2147483648), F1: (("hello" + (v17 + "a")) + v17), F2: -1}
                        var v24 string = ("a" + "\n\t")
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS2{F0: ((3.14159 >= float64(0.1)) || false), F1: 943, F2: v17}
                        v17 = "\n\t"
                        _ = v25
                    }
                    v2 += v2
                    fmt.Println(fi16, fi16)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            case "\n\t":
                v2 = v2
            }
            fmt.Println(fi15)
            _ = fi15
        }
    case "":
        v2 = (v2 >> ((893) & 63))
    }
    v2 = -653
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 float32 = 3.14159
        var v28 uint = 42
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 bool = (!false)
    v29 = (!v29)
    var v30 float64 = (3.14159 - -2.25)
    if v29 {
        v29 = v29
    }
    var v31 int = 2147483647
    for fi32 := 0; fi32 < 13; fi32++ {
        if true {
            delete(v3, 2147483647)
        }
        if (128 == 255) {
            for fi33 := 0; fi33 < 18; fi33++ {
                if len(v1) > 0 {
                    v1[0] = 2
                }
                v31 = (fi32 / ((fi33) | 1))
                _ = fi33
            }
        }
        _ = fi32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
