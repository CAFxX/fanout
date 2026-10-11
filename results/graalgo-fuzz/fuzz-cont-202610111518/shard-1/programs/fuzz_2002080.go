package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 1
    var v3 bool = (!(!true))
    var v4 int16 = 32767
    var v5 string = ("a" + "\n\t")
    v4 = 1000
    var v6 int32 = -1
    v7 := []int{(^(v2)), 32768}
    v3 = v3
    fmt.Println(v6)
    v8 := GenS1{F0: 0}
    var v9 uint64 = 65
    for fi10 := 0; fi10 < 14; fi10++ {
        delete(v1, (fi10 & (fi10 ^ (-(fi10)))))
        _ = fi10
    }
    v11 := GenS1{F0: 372}
    switch (^(v2)) {
    case 32767:
        if (v9 > (v9 * 753)) {
            var v12 uint64 = 2
            fmt.Println(v2, v6)
            var v13 int16 = 1000
            _ = v12
            _ = v13
        } else {
            switch (v2 >> ((762) & 63)) {
            case 128:
                switch v2 {
                case 2147483647:
                    for fi14 := 0; fi14 < 19; fi14++ {
                        v15 := GenS2{F0: -1}
                        var v16 uint64 = v9
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                case 135:
                    if (v3 || (42 > 1)) {
                        v5 = (v5 + "a")
                        fmt.Println(v6)
                        fmt.Println(v4)
                    } else {
                        v9 += (^(v9))
                    }
                case 343:
                    var v17 int16 = -32768
                    _ = v17
                }
            case 938:
                for fi18 := 0; fi18 < 13; fi18++ {
                    if (v3 && (-622 != v4)) {
                        v2 = v2
                        var v19 float64 = float64(3.14159)
                        switch ("a" + ("a" + v5)) {
                        case "\n\t":
                            v6 = v6
                        case "a":
                            v20 := GenS2{F0: -1}
                            _ = v20
                        case "hello":
                            var v21 uint = 799
                            _ = v21
                        default:
                            v22 := GenS1{F0: 0}
                            _ = v22
                        }
                        _ = v19
                    }
                    _ = fi18
                }
            case 127:
                var v23 int64 = 1
                _ = v23
            }
            for fi24 := 0; fi24 < 7; fi24++ {
                v4 = (v4 << ((4294967295) & 63))
                for fi25 := 0; fi25 < 5; fi25++ {
                    fmt.Println(v6, v6, fi24)
                    _ = fi25
                }
                var v26 uint32 = 1
                _ = fi24
                _ = v26
            }
        }
    }
    var v27 uint16 = 255
    v5 = v5
    fmt.Println(v3, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
