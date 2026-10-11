package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int = 4294967295
    var v2 float64 = 3.14159
    var v3 uint64 = 42
    var v4 uint16 = 1
    v5 := make(map[int]string)
    v6 := make(map[int]bool)
    fmt.Println(v2, v2, v4)
    switch "\n\t" {
    case "hello":
        var v7 int64 = int64(v4)
        _ = v7
    case "foo bar":
        v8 := GenS1{F0: 58}
        _ = v8
    case "a":
        var v9 uint32 = 2147483648
        _ = v9
    }
    var v10 int64 = -804
    v6[int(v3)] = true
    fmt.Println(v3)
    v11 := []int{(v1 >> ((255) & 63))}
    if (395 == 4294967295) {
        switch (v1 / (((v1 & -668)) | 1)) {
        case 32768:
            var v12 uint8 = 138
            _ = v12
        case 443:
            if false {
                if (0.1 <= float32(-2.25)) {
                    v13 := GenS1{F0: (v10 & 42)}
                    _ = v13
                }
                if true {
                    _ = len(v11)
                    v3 = 42
                    v14 := GenS1{F0: (v10 | v10)}
                    _ = v14
                } else {
                    var v15 uint = 0
                    v1 = -258
                    _ = v15
                }
            }
        case 1:
            var v16 float64 = 1.5
            _ = v16
        }
    } else {
        v17 := GenS1{F0: v10}
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 string = "foo bar"
            var v21 float32 = (0.1 / 0.1)
            v11 = append(v11, fi18)
            _ = fi19
            _ = v20
            _ = v21
        }
        if false {
            _ = len(v5)
        } else {
            for fi22 := 0; fi22 < 14; fi22++ {
                if ((v10 % ((v10) | 1)) != v10) {
                    var v23 uint = 423
                    var v24 uint = v23
                    var v25 uint = 1
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    if false {
                        switch "\n\t" {
                        case "foo bar":
                            v2 -= 1.5
                        }
                        var v26 uint8 = 255
                        var v27 string = "\n\t"
                        _ = v26
                        _ = v27
                    }
                    v4 = v4
                }
                v2 = ((3.14159 + v2) * (3.14159 * (1e10 + v2)))
                _ = fi22
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 string = "\n\t"
                _ = fi28
                _ = v29
            }
        }
        v5[(fi18 >> ((uint(v3)) & 63))] = "hello"
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
