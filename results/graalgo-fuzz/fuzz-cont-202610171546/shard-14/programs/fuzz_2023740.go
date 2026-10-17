package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 1
    var v3 uint = 597
    if true {
        _ = len(v1)
    } else {
        v1[-905] = ((("a" + ("\n\t" + "a")) + "hello") + "a")
    }
    if false {
        v2 -= -831
        v3 = (v3 & (v3 ^ v3))
    } else {
        if true {
            if true {
                var v4 int64 = 42
                var v5 uint16 = 120
                for fi6 := 0; fi6 < 11; fi6++ {
                    switch 372 {
                    case 0:
                        var v7 bool = false
                        _ = v7
                    }
                    _ = fi6
                }
                _ = v4
                _ = v5
            }
            fmt.Println(v3, v2)
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 uint32 = 65535
                _ = fi8
                _ = v9
            }
        }
    }
    if false {
        if true {
            fmt.Println(v3)
            v3 = (v3 & v3)
            var v10 int32 = 854
            _ = v10
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                if false {
                    if false {
                        fmt.Println(v2, v3, v3)
                        switch "foo bar" {
                        case "a":
                            switch "foo bar" {
                            case "":
                                for fi12 := 0; fi12 < 13; fi12++ {
                                    if false {
                                        v13 := GenS1{F0: 42}
                                        _ = v13
                                    }
                                    v2 = 397
                                    var v14 int32 = -2147483648
                                    _ = fi12
                                    _ = v14
                                }
                            case "foo bar":
                                v3 = v3
                            case "\n\t":
                                v15 := []int{fi11, (fi11 >> ((v3) & 63)), fi11}
                                _ = v15
                            }
                        }
                        v1[(^(fi11))] = "\n\t"
                    } else {
                        var v16 uint8 = 127
                        for fi17 := 0; fi17 < 2; fi17++ {
                            v2 -= (^(v2))
                            v1[-982] = "\n\t"
                            _ = fi17
                        }
                        v16 = v16
                        _ = v16
                    }
                    for fi18 := 0; fi18 < 19; fi18++ {
                        fmt.Println(fi11, v2, v3)
                        v3 += 42
                        _ = fi18
                    }
                    v2 = 42
                } else {
                    fmt.Println(fi11, v2, v3)
                    v2 += (v2 * v2)
                }
                v19 := []int{}
                v3 += 4294967295
                _ = fi11
                _ = v19
            }
        }
    } else {
        var v20 uint8 = 0
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v2 = v2
        _ = fi21
    }
    if false {
        if (32767 <= 42) {
            if false {
                var v22 int32 = 910
                _ = v22
            } else {
                v23 := []int{32768, 41}
                _ = v23
            }
        }
        if false {
            v24 := []int{2147483647, 901, -550, -351}
            v3 -= v3
            v25 := GenS1{F0: 803}
            _ = v24
            _ = v25
        }
        v26 := GenS1{F0: 2147483648}
        _ = v26
    }
    fmt.Println(v2, v3)
    var v27 uint64 = 18446744073709551615
    v28 := GenS1{F0: 42}
    var v29 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
