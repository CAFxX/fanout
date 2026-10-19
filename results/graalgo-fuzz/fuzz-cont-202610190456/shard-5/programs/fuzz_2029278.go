package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint32) uint64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = (^(p0))
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(fi7, fi7)
            switch ("\n\t" + "") {
            case "hello":
                p0 = 811
            case "\n\t":
                if true {
                    fmt.Println(p0, fi8, fi7)
                } else {
                    var v9 uint = 743
                    _ = v9
                }
            }
            _ = fi8
        }
        _ = fi7
    }
    v10 := []int{}
    v11 := GenS1{F0: (^(p0)), F1: 0}
    var v12 int64 = -1
    var v13 uint64 = 42
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 uint32 = 781
    var v2 uint = 255
    v3 := []string{"\n\t", "hello", ("" + "a"), "\n\t"}
    var v4 float64 = 1.5
    var v5 int = 611
    v6 := make(map[string]bool)
    var v14 uint = 2
    var v15 uint8 = 1
    var v16 int16 = 1000
    v5 += v5
    _ = len(v6)
    for fi17 := 0; fi17 < 19; fi17++ {
        if (150 <= 100) {
            if (42 == genFn0((^(v1)))) {
                var v18 string = ""
                _ = v18
            }
            var v19 string = "foo bar"
            _ = v19
        }
        fmt.Println(v16, v4)
        v6[(("\n\t" + "hello") + "\n\t")] = true
        _ = fi17
    }
    if false {
        v16 -= int16(v1)
        if (false && (false && (!true))) {
            v20 := GenS1{F0: v1, F1: 119}
            if (-2147483648 >= 2147483647) {
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 string = "\n\t"
                    _ = fi21
                    _ = v22
                }
                fmt.Println(v15, v4, v4)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v20
        }
    }
    var v23 int = (v5 ^ (v5 + int(v15)))
    if true {
        switch v23 {
        case 481:
            var v24 bool = false
            _ = v24
        default:
            for fi25 := 0; fi25 < 6; fi25++ {
                delete(v6, (((("a" + "\n\t") + "") + "\n\t") + "\n\t"))
                var v26 int32 = 548
                v6["hello"] = false
                _ = fi25
                _ = v26
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
