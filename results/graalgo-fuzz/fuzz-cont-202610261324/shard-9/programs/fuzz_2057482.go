package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() string {
    var v6 int32 = 1
    v7 := []int{-9223372036854775808, -539}
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 uint8 = 128
    var v4 int64 = 9223372036854775807
    v5 := make(map[string]bool)
    var v8 int8 = 42
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int64 = v4
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 uint16 = 1
    v3 -= (v3 + (v3 - (v3 & 255)))
    if ("a" <= "") {
        var v13 bool = (100 > v8)
        fmt.Println(v3, v12)
        v14 := GenS1{F0: (v8 * v8)}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        if (492 != 969) {
            if true {
                v16 := []int{(fi15 + (fi15 - -9223372036854775808)), fi15, fi15, fi15}
                v3 += (v3 >> ((223) & 63))
                _ = v16
            } else {
                v17 := GenS1{F0: (^(v8))}
                var v18 uint = 180
                v18 += v18
                _ = v17
                _ = v18
            }
            switch ("\n\t" + ("hello" + ("foo bar" + ""))) {
            case "":
                v19 := []int{fi15, 127, (fi15 << ((472) & 63)), int(v3)}
                _ = v19
            case "hello":
                v20 := GenS1{F0: v8}
                _ = v20
            case "foo bar":
                switch (^(fi15)) {
                case -406:
                    for fi21 := 0; fi21 < 7; fi21++ {
                        var v22 uint8 = v3
                        v4 = (v4 ^ (^(v4)))
                        _ = fi21
                        _ = v22
                    }
                default:
                    if true {
                        v23 := []int{(fi15 - fi15), (fi15 / ((fi15) | 1)), fi15, (-(fi15))}
                        v24 := v5["hello"]
                        v12 -= v12
                        _ = v23
                        _ = v24
                    }
                }
            }
            v4 += (v4 * v4)
        } else {
            if ((((-2147483648 <= (fi15 << ((42) & 63))) || false) || false) && false) {
                v8 = (v8 | (v8 - (v8 & v8)))
                for fi25 := 0; fi25 < 3; fi25++ {
                    var v26 int = int(v12)
                    v8 -= v8
                    _ = fi25
                    _ = v26
                }
            } else {
                var v27 bool = ((v8 * (v8 & (v8 & v8))) != v8)
                for fi28 := 0; fi28 < 18; fi28++ {
                    v29 := []int{int(v12), fi15, fi28}
                    v8 -= 42
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
            var v30 float32 = 1.5
            _ = v30
        }
        v4 = v4
        _ = fi15
    }
    v8 = (v8 % ((v8) | 1))
    v4 = v4
    fmt.Println(v12)
    var v31 uint64 = 1
    var v32 int8 = v8
    switch 482 {
    case -267:
        v12 += 1
    case 0:
        for fi33 := 0; fi33 < 1; fi33++ {
            for fi34 := 0; fi34 < 1; fi34++ {
                v35 := GenS1{F0: v32}
                _ = fi34
                _ = v35
            }
            switch "hello" {
            case "hello":
                fmt.Println(v3)
            case "\n\t":
                switch fi33 {
                case 9223372036854775807:
                    var v36 float32 = (float32((0.1 + 0.1)) / (1e10 - 3.14159))
                    _ = v36
                default:
                    fmt.Println(v31)
                }
            case "foo bar":
                var v37 float64 = 3.14159
                _ = v37
            }
            v31 = 42
            _ = fi33
        }
    case -2147483648:
        var v38 int16 = 1000
        _ = v38
    default:
        v4 = (v4 / ((-806) | 1))
    }
    v3 = (^(v3))
    var v39 string = "hello"
    _ = len(v2)
    v40 := GenS1{F0: v32}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v8
}
