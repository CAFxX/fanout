package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 18446744073709551615
    var v3 int64 = 1
    for fi4 := 0; fi4 < 16; fi4++ {
        switch "" {
        case "hello":
            v5 := GenS1{F0: "hello", F1: -1, F2: 42}
            _ = v5
        case "foo bar":
            fmt.Println(fi4, fi4, v3)
        case "":
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := GenS1{F0: "foo bar", F1: 32767, F2: 319}
                delete(v1, (((("\n\t" + "hello") + ("foo bar" + ("hello" + ""))) + "") + (("a" + "foo bar") + ("\n\t" + "\n\t"))))
                _ = fi6
                _ = v7
            }
        default:
            switch ("\n\t" + "") {
            case "hello":
                for fi8 := 0; fi8 < 12; fi8++ {
                    v9 := GenS1{F0: ("a" + ((("" + ("hello" + "hello")) + "hello") + "foo bar")), F1: 1, F2: 42}
                    _ = fi8
                    _ = v9
                }
            case "a":
                switch (("hello" + ("foo bar" + "hello")) + "foo bar") {
                case "a":
                    switch "a" {
                    case "foo bar":
                        var v10 float32 = 3.14159
                        _ = v10
                    default:
                        v3 = v3
                    }
                }
            }
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            switch fi11 {
            case 2147483647:
                if (33 >= 1) {
                    fmt.Println(fi4)
                    v12 := GenS1{F0: "hello", F1: 32767, F2: 255}
                    _ = v12
                } else {
                    if true {
                        v2 += v2
                        switch 410 {
                        case -1:
                            v1["hello"] = 1
                        case 32767:
                            var v13 float64 = (0.1 + (float64(1.5) / float64(-2.25)))
                            _ = v13
                        case 594:
                            v14 := GenS1{F0: "hello", F1: 47, F2: 1}
                            _ = v14
                        }
                    }
                    var v15 uint16 = 32767
                    _ = v15
                }
            case 4294967295:
                v2 += 42
            case 388:
                v2 -= (v2 * (v2 - v2))
            }
            var v16 int64 = v3
            v16 = (v16 << ((0) & 63))
            _ = fi11
            _ = v16
        }
        _ = fi4
    }
    v2 = (v2 & 942)
    for fi17 := 0; fi17 < 18; fi17++ {
        v2 -= (^(v2))
        fmt.Println(fi17, v3, fi17)
        _ = fi17
    }
    v18 := []int{-2147483648}
    fmt.Println(v3, v3)
    if (true && true) {
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := v1["foo bar"]
            _ = fi19
            _ = v20
        }
        var v21 float32 = 0.1
        v2 = v2
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v2 -= 1
        _ = fi22
    }
    v2 = v2
    fmt.Println(v2, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
}
