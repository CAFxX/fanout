package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 uint16) uint64 {
    p0 = (^(p0))
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 483}
        var v8 uint32 = 2
        v9 := []int{fi6, fi6, fi6, fi6}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := []string{("" + ("hello" + ("foo bar" + (("" + "\n\t") + ("foo bar" + "\n\t")))))}
    v2 := []int{}
    var v3 int8 = 17
    var v4 int = 32767
    var v5 uint32 = 1
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int = 128
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v12 int64 = -1
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint8 = 42
    v5 = (^(v5))
    if false {
        fmt.Println(v13, v3, v4)
        for fi14 := 0; fi14 < 9; fi14++ {
            v4 = 128
            if false {
                if (1 == 0) {
                    if false {
                        switch (v4 + fi14) {
                        case 32768:
                            var v15 uint64 = 484
                            _ = v15
                        case 667:
                            v5 = v5
                        }
                        for fi16 := 0; fi16 < 1; fi16++ {
                            v13 = 204
                            _ = fi16
                        }
                        v3 = (v3 & -76)
                    }
                    var v17 int = (v4 >> ((255) & 63))
                    v18 := GenS1{F0: v5}
                    _ = v17
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        switch (v4 | fi19) {
                        case 0:
                            switch (v4 % ((123) | 1)) {
                            case 922:
                                var v20 int8 = -1
                                _ = v20
                            }
                        }
                        var v21 uint16 = 1
                        _ = fi19
                        _ = v21
                    }
                    if len(v1) > 0 {
                        v1[0] = ("\n\t" + "foo bar")
                    }
                }
            }
            switch ("a" + "hello") {
            case "foo bar":
                v3 = v3
            }
            _ = fi14
        }
        v22 := GenS2{F0: v3, F1: v4, F2: 2}
        _ = v22
    }
    v3 -= 100
    if true {
        if true {
            v13 += 111
        }
    } else {
        v13 -= 255
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        if (42 <= 625) {
            var v24 float64 = 0.0
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v24
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            fmt.Println(fi23, fi25, fi23)
            _ = fi25
        }
        _ = fi23
    }
    switch v4 {
    case -9223372036854775808:
        fmt.Println(v13)
    case 128:
        v26 := GenS3{F0: (v4 & -9223372036854775808)}
        _ = v26
    }
    fmt.Println(v4, v4, v4)
    var v27 int = (v4 & (v4 - v4))
    v5 += v5
    fmt.Println(v4, v27, v3)
    var v28 int16 = 32767
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
