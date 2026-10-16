package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
func genFn0() int8 {
    v5 := GenS1{F0: 1, F1: 1}
    switch "" {
    case "hello":
        if false {
            for fi6 := 0; fi6 < 16; fi6++ {
                fmt.Println(fi6)
                _ = fi6
            }
            var v7 int8 = 100
            _ = v7
        }
    }
    var v8 string = ""
    _ = v5
    _ = v8
    return 127
}

func main() {
    v1 := []int{42, 127}
    var v2 int64 = -1
    var v3 int8 = 80
    v4 := make(map[int]string)
    v2 -= (v2 % ((v2) | 1))
    switch 9223372036854775807 {
    case 128:
        if false {
            var v9 int64 = (v2 << ((42) & 63))
            var v10 int64 = v9
            _ = v10
            _ = v9
        }
    case 687:
        var v11 bool = false
        _ = v11
    }
    fmt.Println(v3, v2)
    for fi12 := 0; fi12 < 2; fi12++ {
        if (2 != 90) {
            v2 = v2
        }
        _ = fi12
    }
    var v13 uint16 = 663
    fmt.Println(v3, v2)
    v14 := GenS2{F0: 1}
    if true {
        switch -1 {
        case -913:
            switch (("foo bar" + "\n\t") + "a") {
            case "":
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 uint64 = 2
                    v17 := GenS2{F0: (^(v16))}
                    fmt.Println(fi15, v3, v2)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
        }
        delete(v4, 9223372036854775807)
        for fi18 := 0; fi18 < 16; fi18++ {
            v13 += v13
            _ = len(v4)
            if (12 != 1) {
                v19 := GenS3{F0: (1 <= 0)}
                var v20 uint64 = 2
                v3 = v3
                _ = v19
                _ = v20
            }
            _ = fi18
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
