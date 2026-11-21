package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            if (-1 < -515) {
                var v6 uint64 = 105
                _ = v6
            } else {
                fmt.Println(fi5)
            }
            v1 = (v1 + "foo bar")
            v7 := GenS1{F0: 777, F1: 65535, F2: (("foo bar" + (v1 + v1)) + v1)}
            _ = fi5
            _ = v7
        }
        for fi8 := 0; fi8 < 5; fi8++ {
            v1 = v1
            fmt.Println(fi4, fi4, v1)
            _ = fi8
        }
        var v9 uint64 = 502
        _ = fi4
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        switch fi10 {
        case 30:
            fmt.Println(fi10, v1)
        default:
            switch "a" {
            case "a":
                fmt.Println(v1, fi10)
            case "\n\t":
                switch (fi10 * 1) {
                case 884:
                    var v11 float32 = 0.0
                    _ = v11
                case -807:
                    v12 := GenS2{F0: 32767, F1: -730}
                    _ = v12
                }
            case "hello":
                v1 = ("\n\t" + v1)
            default:
                for fi13 := 0; fi13 < 8; fi13++ {
                    if false {
                        v14 := GenS2{F0: 399, F1: 42}
                        _ = v14
                    } else {
                        delete(v3, fi13)
                    }
                    v15 := GenS3{F0: 9223372036854775807}
                    _ = fi13
                    _ = v15
                }
            }
        }
        _ = fi10
    }
    v16 := GenS1{F0: 94, F1: 18446744073709551615, F2: (("hello" + (v1 + ("\n\t" + v1))) + v1)}
    var v17 string = v1
    v17 = ("\n\t" + v17)
    if (42 == 18446744073709551615) {
        if true {
            v18 := []int{94, 127}
            var v19 uint = 4294967295
            _ = v18
            _ = v19
        } else {
            var v20 int16 = 1
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := GenS1{F0: 2, F1: 913, F2: ((v1 + v17) + v1)}
        v1 = v1
        switch (fi21 ^ 9223372036854775807) {
        case -1:
            v1 = (v1 + "")
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v17)
    var v23 uint64 = 2
    switch v17 {
    case "a":
        if true {
            for fi24 := 0; fi24 < 20; fi24++ {
                if true {
                    fmt.Println(v23, v17, v1)
                    for fi25 := 0; fi25 < 12; fi25++ {
                        if true {
                            _ = len(v3)
                            var v26 uint = 65535
                            _ = v26
                        } else {
                            var v27 float32 = 0.0
                            v28 := []int{fi25, (fi25 % ((fi25) | 1))}
                            v1 = v17
                            _ = v27
                            _ = v28
                        }
                        var v29 string = v17
                        _ = fi25
                        _ = v29
                    }
                    v17 = v17
                }
                v30 := v2[807]
                var v31 int = v30
                _ = fi24
                _ = v30
                _ = v31
            }
            var v32 string = v1
            var v33 int64 = 553
            _ = v32
            _ = v33
        } else {
            v23 += v23
            fmt.Println(v17)
            var v34 float32 = 3.14159
            _ = v34
        }
    case "hello":
        v35 := GenS4{F0: v1}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
}
