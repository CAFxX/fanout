package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 string, p1 int) int8 {
    fmt.Println(p0)
    if (!false) {
        if false {
            p0 = "a"
            if true {
                fmt.Println(p0)
            } else {
                p0 = p0
                fmt.Println(p0, p1)
                if (((1e10 / -2.25) != 1e10) && ((!(-780 > 42)) || (false && ((true && true) && (811 <= 670))))) {
                    if true {
                        for fi4 := 0; fi4 < 1; fi4++ {
                            p1 += (fi4 - 226)
                            p1 = (fi4 << ((857) & 63))
                            _ = fi4
                        }
                        switch "hello" {
                        case "a":
                            v5 := GenS1{F0: 32767, F1: 42}
                            _ = v5
                        case "hello":
                            var v6 float64 = 3.14159
                            _ = v6
                        case "":
                            var v7 float32 = (0.0 * -2.25)
                            _ = v7
                        }
                        var v8 int32 = 649
                        _ = v8
                    } else {
                        v9 := GenS1{F0: 574, F1: -553}
                        var v10 int32 = -1
                        if true {
                            v11 := []int{}
                            p0 = ("a" + "foo bar")
                            fmt.Println(v10, p0, p1)
                            _ = v11
                        }
                        _ = v10
                        _ = v9
                    }
                } else {
                    p0 = ("hello" + p0)
                    switch (p1 + -720) {
                    case -2147483648:
                        fmt.Println(p0)
                    }
                }
            }
        }
    } else {
        if ((-2.25 < -2.25) || true) {
            fmt.Println(p1, p0)
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 2
    var v3 string = "hello"
    v12 := GenS1{F0: 32767, F1: 1}
    v13 := []int{-1, 0, 1}
    fmt.Println(v3)
    var v14 int64 = 0
    v3 = "hello"
    v15 := v1[""]
    v2 = v2
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v13 = append(v13, 2147483647)
    v16 := GenS1{F0: 255, F1: -234}
    if (-1 >= -487) {
        for fi17 := 0; fi17 < 17; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(v2, v14, v2)
                _ = fi18
            }
            fmt.Println(fi17, fi17)
            _ = fi17
        }
        if (1.5 < 0.1) {
            var v19 int8 = 1
            _ = v19
        } else {
            fmt.Println(v3)
            switch 32768 {
            case 42:
                var v20 uint8 = (v2 ^ (v2 | v2))
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            for fi22 := 0; fi22 < 16; fi22++ {
                switch (v15 + ("hello" + (v3 + v3))) {
                case "foo bar":
                    v23 := GenS1{F0: 663, F1: 1000}
                    v15 = v15
                    _ = v23
                default:
                    v15 = v3
                }
                v14 = 9223372036854775807
                v24 := GenS1{F0: 0, F1: 32767}
                _ = fi22
                _ = v24
            }
            v3 = (v3 + v15)
            delete(v1, (v3 + v15))
            _ = fi21
        }
    }
    v25 := GenS2{F0: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
}
