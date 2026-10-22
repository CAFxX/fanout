package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
func genFn0() bool {
    if (-2.25 == float32(1.5)) {
        if true {
            v5 := []int{-9223372036854775808, -25}
            v5 = append(v5, -2147483648)
            var v6 uint = 255
            _ = v5
            _ = v6
        }
        var v7 uint16 = 255
        _ = v7
    } else {
        var v8 int = 128
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        switch fi9 {
        case 0:
            fmt.Println(fi9)
        case 4294967295:
            if true {
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS1{F0: 4294967295}
                    var v12 string = (("\n\t" + "a") + "foo bar")
                    fmt.Println(fi10)
                    switch (v12 + (("foo bar" + v12) + ((v12 + v12) + "\n\t"))) {
                    case "":
                        v13 := GenS2{F0: 255}
                        _ = v13
                    case "\n\t":
                        switch (fi10 >> ((413) & 63)) {
                        case 32768:
                            v14 := GenS2{F0: 2}
                            v15 := []int{fi10, fi9, 817}
                            v12 = v12
                            _ = v14
                            _ = v15
                        }
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v16 int32 = 1
                _ = v16
            } else {
                switch ("foo bar" + "foo bar") {
                case "a":
                    v17 := GenS2{F0: 1}
                    _ = v17
                case "\n\t":
                    if true {
                        for fi18 := 0; fi18 < 14; fi18++ {
                            var v19 float64 = 1.5
                            v20 := []int{1, fi18}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        if ((-587 < 42) && (255 >= 65535)) {
                            fmt.Println(fi9)
                        } else {
                            v21 := []int{-2147483648, (fi9 | fi9), -9223372036854775808, (fi9 << ((413) & 63))}
                            _ = v21
                        }
                        var v22 int32 = -2147483648
                        _ = v22
                    }
                }
            }
        }
        _ = fi9
    }
    var v23 bool = (982 < 65535)
    for fi24 := 0; fi24 < 12; fi24++ {
        v23 = (2147483647 <= 2)
        v25 := []int{}
        _ = fi24
        _ = v25
    }
    var v26 bool = (-554 <= 0)
    _ = v23
    _ = v26
    return (1 <= 206)
}
func genFn1() int {
    var v27 uint = 18446744073709551615
    v27 = v27
    var v28 int = 2147483647
    _ = v27
    _ = v28
    return -486
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 int16 = -32768
    v29 := GenS2{F0: 165}
    v1 -= 1000
    var v30 int = genFn1()
    var v31 float32 = 0.1
    var v32 uint16 = 0
    v4 += 42
    for fi33 := 0; fi33 < 15; fi33++ {
        fmt.Println(v31, v32, v30)
        _ = fi33
    }
    var v34 int8 = -1
    v35 := GenS2{F0: v32}
    v34 += (v34 * (^(v34)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
}
