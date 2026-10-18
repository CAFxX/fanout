package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int32 }
func genFn0(p0 int8, p1 uint32) int {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int8 = p0
        _ = fi5
        _ = v6
    }
    v7 := []int{317, 905, -1, 32768}
    switch 0 {
    case -156:
        var v8 int32 = -2147483648
        _ = v8
    case -226:
        var v9 uint = 2
        _ = v9
    case 4294967295:
        var v10 bool = (2147483647 >= 42)
        _ = v10
    }
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}

func main() {
    var v1 int = -886
    var v2 int64 = 891
    var v3 int8 = 1
    var v4 bool = (127 <= 235)
    var v11 uint = 348
    fmt.Println(v11, v11, v2)
    if (!true) {
        for fi12 := 0; fi12 < 16; fi12++ {
            v11 = v11
            v11 += (^(v11))
            v13 := GenS1{F0: 42}
            _ = fi12
            _ = v13
        }
        switch ("\n\t" + "") {
        case "hello":
            switch 42 {
            case 2147483647:
                fmt.Println(v2, v1, v11)
            case -1:
                var v14 uint32 = 42
                _ = v14
            case 790:
                v4 = (false && true)
            }
        case "\n\t":
            var v15 uint = (v11 / ((65535) | 1))
            _ = v15
        }
        v16 := []int{}
        _ = v16
    }
    v17 := []int{genFn0(v3, 42), 42}
    var v18 int = (^(v1))
    v17 = append(v17, (v18 * -547))
    var v19 string = ("" + "hello")
    var v20 float32 = 0.0
    var v21 int16 = 42
    if len(v17) > 0 {
        v17[0] = int(v2)
    }
    if (!v4) {
        if v4 {
            v19 = (((v19 + (("\n\t" + v19) + ("" + "\n\t"))) + (v19 + v19)) + (("" + ((v19 + ("hello" + v19)) + v19)) + "\n\t"))
            for fi22 := 0; fi22 < 15; fi22++ {
                switch v19 {
                case "":
                    v23 := GenS1{F0: (v21 >> ((255) & 63))}
                    _ = v23
                case "\n\t":
                    v20 = v20
                case "a":
                    for fi24 := 0; fi24 < 8; fi24++ {
                        var v25 bool = (!v4)
                        _ = fi24
                        _ = v25
                    }
                default:
                    var v26 int64 = v2
                    _ = v26
                }
                switch (v18 + (fi22 * (v18 << (((^(v11))) & 63)))) {
                case 4294967295:
                    if (!false) {
                        v2 += v2
                        var v27 uint64 = 2
                        _ = v27
                    } else {
                        v28 := GenS2{F0: genFn0((v3 & int8(v11)), uint32(v11)), F1: 2147483647}
                        fmt.Println(v4, v20)
                        _ = v28
                    }
                }
                _ = fi22
            }
            fmt.Println(v4, v18, v2)
        } else {
            v18 -= 127
            switch genFn0(v3, 2147483647) {
            case -149:
                for fi29 := 0; fi29 < 18; fi29++ {
                    v3 -= v3
                    fmt.Println(v2, v3)
                    _ = fi29
                }
            case -327:
                v30 := GenS1{F0: (v21 * v21)}
                _ = v30
            case 0:
                if (v21 != 1) {
                    var v31 uint32 = 188
                    _ = v31
                } else {
                    v20 += v20
                    var v32 uint64 = 1
                    var v33 uint8 = 24
                    _ = v32
                    _ = v33
                }
            }
        }
    } else {
        var v34 string = (v19 + v19)
        _ = v34
    }
    var v35 uint8 = 1
    if true {
        if (164 == 2147483648) {
            v3 += 100
            fmt.Println(v1, v35, v20)
        }
    }
    fmt.Println(v11, v3)
    fmt.Println(v18, v4, v3)
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    var v36 int64 = -1
    if len(v17) > 0 {
        v17[0] = (v18 - -668)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
