package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int8) int16 {
    fmt.Println(p0)
    switch ("a" + "\n\t") {
    case "\n\t":
        var v6 string = ("a" + "a")
        _ = v6
    case "foo bar":
        if (0 <= 440) {
            var v7 float64 = 0.1
            if (255 != 255) {
                var v8 uint16 = 2
                _ = v8
            }
            _ = v7
        }
    }
    v9 := GenS1{F0: -577}
    if true {
        var v10 int64 = 0
        if (173 < -1) {
            if true {
                fmt.Println(p0, p0)
                p0 = p0
                fmt.Println(p0, v10)
            } else {
                if ((-(v10)) <= v10) {
                    p0 = p0
                    var v11 int = -1
                    _ = v11
                }
            }
            switch "a" {
            case "hello":
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 uint16 = 0
                    v13 += v13
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 int32 = 1
                        _ = fi14
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
            case "\n\t":
                var v16 int = -397
                _ = v16
            }
            if false {
                v17 := []int{997}
                _ = v17
            }
        } else {
            var v18 bool = false
            _ = v18
        }
        _ = v10
    }
    _ = p0
    _ = v9
    return 553
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 65535
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    var v5 int32 = 0
    var v19 uint8 = 76
    for fi20 := 0; fi20 < 3; fi20++ {
        v19 += v19
        _ = fi20
    }
    v21 := GenS1{F0: v5}
    v19 = (^(v19))
    v19 = 1
    v22 := GenS1{F0: v5}
    var v23 uint32 = 815
    if (!(!(100 < 42))) {
        v23 = v23
        v19 = v19
        fmt.Println(v2, v19, v5)
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        var v24 float32 = -2.25
        _ = v24
    case "hello":
        var v25 uint32 = (v23 | (v23 - (^(v23))))
        _ = v25
    case "":
        v26 := GenS1{F0: v5}
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
