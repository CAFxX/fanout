package main

import (
    "fmt"
)

func genFn0(p0 uint8, p1 int16, p2 int16) uint32 {
    p0 += (p0 << ((0) & 63))
    p2 += p2
    v6 := []int{2147483647, 0, -115}
    if false {
        if (true || false) {
            if len(v6) > 0 {
                v6[0] = 283
            }
            var v7 int8 = 42
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 874
}

func main() {
    var v1 uint = 729
    var v2 uint32 = 1
    var v3 int32 = -2147483648
    v4 := make(map[int]bool)
    var v5 string = "a"
    v3 -= v3
    var v8 int16 = 347
    v8 -= v8
    var v9 int16 = 42
    fmt.Println(v5, v2, v2)
    var v10 int8 = 0
    if (true || (true || ((304 == 0) && (651 < 1)))) {
        var v11 int16 = (v8 ^ v9)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 11; fi13++ {
                switch fi12 {
                case 42:
                    if (false || true) {
                        switch 1 {
                        case -9223372036854775808:
                            v14 := []int{}
                            _ = v14
                        }
                        v4[(fi13 * (fi12 / ((fi13) | 1)))] = false
                    } else {
                        v5 = (((v5 + "foo bar") + "\n\t") + "a")
                        var v15 int64 = 9223372036854775807
                        switch v5 {
                        case "hello":
                            if true {
                                var v16 string = (("a" + "foo bar") + (("foo bar" + "") + (v5 + ((((v5 + "foo bar") + "\n\t") + v5) + v5))))
                                _ = len(v4)
                                v3 = (v3 >> ((42) & 63))
                                _ = v16
                            } else {
                                var v17 uint32 = (v2 - 311)
                                _ = v17
                            }
                        }
                        switch (v5 + v5) {
                        case "foo bar":
                            v18 := v4[fi13]
                            _ = v18
                        case "\n\t":
                            v3 += v3
                        }
                        _ = v15
                    }
                default:
                    v19 := []int{(fi12 * fi13), fi13}
                    _ = v19
                }
                _ = len(v4)
                var v20 int64 = 42
                _ = fi13
                _ = v20
            }
            v10 = 85
            _ = fi12
        }
        v8 -= (v9 << (((v1 / ((uint(v2)) | 1))) & 63))
    }
    if true {
        var v21 uint64 = 1
        v8 = (v8 | (^(v9)))
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
